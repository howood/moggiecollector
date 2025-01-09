package custommiddleware

import (
	"github.com/howood/moggiecollector/di/svcluster"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

// XRequestID create X-Request-Id for each Request.
func RequestLog(svCluster *svcluster.ServiceCluster) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			if err := svCluster.RequestLogSV.CreateRequest(ctx, c); err != nil {
				log.Error(ctx, err.Error())
			}

			return next(c)
		}
	}
}
