package exception

import "C"
import "github.com/gin-gonic/gin"

type BusinessException struct {
	Status int         `json:"status"`
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (businessException *BusinessException) Error() string {
	return businessException.Msg
}

func AbortWithStatusMsg(ctx *gin.Context, status int, code int, msg string, data interface{}) {
	result := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	ctx.AbortWithStatusJSON(status, result)
}

func NewBusinessEception(status int, code int, msg string, data interface{}) BusinessException {
	return BusinessException{
		Status:status,
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
