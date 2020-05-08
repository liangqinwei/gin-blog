package goException

import (
	"fmt"
	"ginBlog/src/gin-blog/middleware/goLogger"
	"ginBlog/src/gin-blog/package/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"runtime"
)


func RecoveryMiddleware() gin.HandlerFunc {
	/**
	全屏中间件捕获异常
	 */
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
			    //判断是否为自定义异常
				if e, ok := err.(exception.BusinessException);ok{
					exception.AbortWithStatusMsg(context,e.Status,e.Code,e.Msg,e.Data)
					//打印堆栈
					stack := make([]byte, 1024*8)
					stack = stack[:runtime.Stack(stack, false)]
					httpRequest, _ := httputil.DumpRequest(context.Request, false)
					reqMethod:=context.Request.Method
					reqUrl:=context.Request.RequestURI
					statusCode:=context.Writer.Status()
					clientIp:=context.ClientIP()
					goLogger.Logger.WithFields(logrus.Fields{
						"status_code":statusCode,
						"http_request":string(httpRequest),
						"client_ip":clientIp,
						"req_method":reqMethod,
						"req_url":reqUrl,
						"error":err,
						"stack":fmt.Errorf("%s",stack),
					}).Info()
					return
				}
				//打印堆栈
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				httpRequest, _ := httputil.DumpRequest(context.Request, false)
				reqMethod:=context.Request.Method
				reqUrl:=context.Request.RequestURI
				statusCode:=context.Writer.Status()
				clientIp:=context.ClientIP()
				goLogger.Logger.WithFields(logrus.Fields{
					"status_code":statusCode,
					"http_request":string(httpRequest),
					"client_ip":clientIp,
					"req_method":reqMethod,
					"req_url":reqUrl,
					"error":err,
					"stack":fmt.Errorf("%s",stack),
				}).Info()
				//fmt.Errorf("[Recovery] panic recovered:\n%s\n%s\n%s", string(httprequest), err, stack)
				//todo 异常报警
				exception.AbortWithStatusMsg(context,http.StatusInternalServerError,http.StatusInternalServerError,"",nil)
			}
		}()
		context.Next()
	}
}
