package goLogger

import (
	"bytes"
	"ginBlog/src/gin-blog/package/setting"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	INFO  = "InfoLevel"
	DEBUG = "DebugLevel"
	FATAL = "FatalLevel"
	ERROR = "ErrorLevel"
	WARN  = "WarnLevel"
	TRACE = "TraceLevel"
)

var (Logger *logrus.Logger)

func init(){
	logFilePath := setting.LogPath
	LogFileName := setting.LogName
	//日志文件地址
	fileName := path.Join(logFilePath, LogFileName)
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal("打开日志文件失败:%v", err)
	}

	//实例化
	Logger = logrus.New()
	//设置输出
	Logger.Out = src

	logLevel := setting.LogLevel
	if logLevel == INFO {
		Logger.SetLevel(logrus.InfoLevel)
	}
	if logLevel == DEBUG {
		Logger.SetLevel(logrus.DebugLevel)
	}
	if logLevel == FATAL {
		Logger.SetLevel(logrus.FatalLevel)
	}
	if logLevel == WARN {
		Logger.SetLevel(logrus.WarnLevel)
	}
	if logLevel == TRACE {
		Logger.SetLevel(logrus.TraceLevel)
	}
	if logLevel == ERROR {
		Logger.SetLevel(logrus.ErrorLevel)
	}
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}
	hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2020-01-02 15:04:05"})
	Logger.AddHook(hook)
}


func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data string
		reqMethod:=c.Request.Method
		reqUrl:=c.Request.RequestURI
		statusCode:=c.Writer.Status()
		clientIp:=c.ClientIP()
		if reqMethod == http.MethodPost { // 如果是post请求，则读取body
			body, err := c.GetRawData() // body 只能读一次，读出来之后需要重置下 Body
			if err != nil {
				log.Fatal(err)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 重置body

			data = string(body)
		}
		startTime:=time.Now()
		c.Next()
		endTime:=time.Now()
		latencyTime:=endTime.Sub(startTime)

		Logger.WithFields(logrus.Fields{
			"status_code":statusCode,
			"latency_time":latencyTime,
			"client_ip":clientIp,
			"req_method":reqMethod,
			"req_url":reqUrl,
			"body":data,
		}).Info()

	}

}
