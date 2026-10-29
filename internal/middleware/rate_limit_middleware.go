package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"porto-backend-go/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func RateLimiter(r rate.Limit, b int) gin.HandlerFunc {
	var mu sync.RWMutex
	clients := make(map[string]*client)

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, c := range clients {
				if time.Since(c.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.GetHeader("CF-Connecting-IP")
		if ip == "" {
			if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
				ip = strings.TrimSpace(strings.Split(xff, ",")[0])
			}
		}
		if ip == "" {
			ip = c.ClientIP()
		}

		mu.Lock()
		cl, exists := clients[ip]
		if !exists {
			limiter := rate.NewLimiter(r, b)
			clients[ip] = &client{limiter: limiter, lastSeen: time.Now()}
			cl = clients[ip]
		}
		cl.lastSeen = time.Now()
		mu.Unlock()

		if !cl.limiter.Allow() {
			response.Error(c, http.StatusTooManyRequests, "Too many requests")
			c.Abort()
			return
		}

		c.Next()
	}
}
