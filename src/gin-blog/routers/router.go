package routers

import (
	"ginBlog/src/gin-blog/middleware/goException"
	"ginBlog/src/gin-blog/middleware/goLogger"
	"ginBlog/src/gin-blog/package/setting"
	v1 "ginBlog/src/gin-blog/routers/api/v1"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器，必须引入，如若不引入，则需要自己定义
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // 引入对应数据库引擎
	"github.com/GoAdminGroup/go-admin/modules/language"
	_ "github.com/GoAdminGroup/themes/adminlte" // 引入主题，必须引入，不然报错
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	// 实例化一个GoAdmin引擎对象
	eng := engine.Default()

	// GoAdmin全局配置，也可以写成一个json，通过json引入
	cfg := config.Config{
		// 数据库配置，为一个map，key为连接名，value为对应连接信息
		Databases: config.DatabaseList{
			// 默认数据库连接，名字必须为default
			"default": {
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "liangqinwei",
				Name:       "gin_blog",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin", // 访问网站的前缀
		// Store 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		// 网站语言
		Language: language.CN,
	}

	// 增加配置与插件，使用Use方法挂载到Web框架中
	_ = eng.AddConfig(cfg).Use(r)
	// 这里引入你需要管理的业务表配置
	// 后面会介绍如何使用命令行根据你自己的业务表生成Generators
	// AddGenerators(Generators).

	//引入中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//引入异常捕获中间件
	r.Use(goException.RecoveryMiddleware())
	r.Use(goLogger.LoggerToFile())

	gin.SetMode(setting.RunMode)
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/add/tag", v1.AddTag)
		apiV1.POST("/add/tags", v1.AddTags)
		apiV1.PUT("/update/tags/:id", v1.EditTag)
		apiV1.DELETE("/delete/tags/:id", v1.DeleteTag)
		apiV1.GET("/tags/test", v1.TestTag)

	}
	return r
}
