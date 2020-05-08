package main

import (
	"fmt"
	"ginBlog/src/gin-blog/package/setting"
	"ginBlog/src/gin-blog/routers"
	"net/http"
)

func main() {

	router:=routers.InitRouter()

	s:=&http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
