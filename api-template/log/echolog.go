package log

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
				fields = append(fields, zap.String("request_id", id))
			}

			n := res.Status
			switch {
			case n >= 500:
				log.With(zap.Error(err)).Error("HTTP Response: 5XX ("+strconv.Itoa(n)+")", fields...)
			case n >= 400:
				log.With(zap.Error(err)).Error("HTTP Response: 4XX ("+strconv.Itoa(n)+")", fields...)
			case n >= 300:
				log.Warn("HTTP Response: 3XX ("+strconv.Itoa(n)+")", fields...)
			case n >= 200:
				log.Info("HTTP Response: 2XX ("+strconv.Itoa(n)+")", fields...)
			case n >= 100:
				log.Warn("HTTP Response: 1XX ("+strconv.Itoa(n)+")", fields...)
			default:
				log.Info("HTTP Response: 200", fields...)
			}

			return nil
		}
	}
}
