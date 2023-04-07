package controller

import (
	"biz-c-service/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/server"
)

func CfController() {
	// 测试入口和出口的拦截情况
	//server
	server.Get("front/cf", service.Cf)
	server.Get("bc/:haveRedis/:haveMysql", service.Bc)
	//server.Use(TimeoutInterceptor(time.Second * 2))

	server.Get("front/cf/ok/ok", service.FrontCfOkOk)
	server.Post("front/cf/ok/file", service.FrontCfOkFile)
	server.Get("front/cf/ok/stop", service.FrontCfOkStop)
	server.Get("front/cf/stop/ok", service.FrontCfStopOk)
	server.Get("front/cf/stop/stop", service.FrontCfStopStop)

	server.Get("front/cf/:haveRedis/:haveMysql/:haveEtcd", service.FrontCfStopStop)
}

// TimeoutInterceptor 超时配置
func TimeoutInterceptor(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.JSON(http.StatusRequestTimeout, "request timeout")
				c.Abort()
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
