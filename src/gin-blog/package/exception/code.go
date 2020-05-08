package exception

import "net/http"

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	SQL_EXEC_ERROR                 = 20005
	ERROR_UNKNOW                   = 30000
	NOT_FOUND                      = 10004
	BODY_STRUCT_VALIDATE_ERROR     = 10500
)

//500服务器异常
func ServerError() BusinessException {
	return NewBusinessEception(http.StatusInternalServerError,ERROR, GetMsg(ERROR), nil)
}

func NotFound() BusinessException {
	return NewBusinessEception(http.StatusNotFound,NOT_FOUND, GetMsg(NOT_FOUND), nil)
}

func UnknowError(message string) BusinessException {
	return NewBusinessEception(http.StatusBadRequest,ERROR_UNKNOW, message, nil)
}

func ParamterError(message string) BusinessException {
	return NewBusinessEception(http.StatusOK,INVALID_PARAMS, message, nil)
}

func BodyError(message string,data interface{}) BusinessException {
	return NewBusinessEception(http.StatusBadRequest,BODY_STRUCT_VALIDATE_ERROR, message, data)
}

func SqlError(message string) BusinessException{
	return NewBusinessEception(http.StatusOK,SQL_EXEC_ERROR,message,nil)
}


