package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/tuhalang/authen/internal/logger"
)

type stackTrace struct {
	traceID string
}

const (
	traceIDKey = "trace_id"
)

// LoggingMiddleware interceptor handle writing log
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log := logger.Get()

		traceID := xid.New().String()

		ctx := context.WithValue(c.Request.Context(), traceIDKey, stackTrace{traceID: traceID})

		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str(traceIDKey, traceID)
		})

		c.Request = c.Request.WithContext(log.WithContext(ctx))

		log.
			Info().
			Str("method", c.Request.Method).
			Str("url", c.Request.URL.RequestURI()).
			Str("ip", c.Request.RemoteAddr).
			Str("user_agent", c.Request.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Msg("")

		c.Writer.Header().Set("TRACE_ID", traceID)
		c.Next()
	}
}
