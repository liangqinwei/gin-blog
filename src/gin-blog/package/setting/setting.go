package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
	LogPath string
	LogName string
	LogLevel string
)

func init() {
	var err error
	Cfg, err = ini.Load("src/gin-blog/conf/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'gin-blog/conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadLog()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8005)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(30)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(30)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Fail to get section 'app':%v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)

}

func LoadLog() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		log.Fatal("Fail to get Section 'app:%v", err)
	}
	LogPath=sec.Key("LOG_FILE_PATH").MustString("D:\\goland\\workspace\\ginBlog\\src\\log")
	LogName=sec.Key("LOG_FILE_NAME").MustString("system.log")
	LogLevel=sec.Key("LOG_LEVEL").MustString("InfoLevel")

}
