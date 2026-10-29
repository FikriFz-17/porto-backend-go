package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"porto-backend-go/pkg/response"
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
		ip := c.ClientIP()

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
