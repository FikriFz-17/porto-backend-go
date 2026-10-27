# Stage 1: Build binary
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compile binary api & migrate
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/migrate ./cmd/migrate

# Stage 2: Minimal runtime image
FROM alpine:3.20

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/api .
COPY --from=builder /app/migrate .
COPY --from=builder /app/database/migrations ./database/migrations

EXPOSE 8080

CMD ["sh", "-c", "./migrate up && ./api"]