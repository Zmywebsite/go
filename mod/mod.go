package mod

import (
	"net/http"
	"vi-server/core/res"
	"vi-server/core/router"
	"vi-server/mod/user"
	"vi-server/utils"

	_ "vi-server/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r = gin.Default() // 路由方法

func Setupr() *gin.Engine {

	// 欢迎页面
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "欢迎进入vios后端")
	})

	// 静态资源
	r.Static("/public", "./public")

	// 让Swagger与gin结合
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 检查 X-Access-Token 判断是否有请求权限
	r.Use(func(ctx *gin.Context) {

		path := ctx.Request.URL.Path
		println("访问的路径:", path)

		xToken := ctx.Request.Header.Get("X-Access-Token") // 获取请求头

		if xToken != utils.GetConfig().App.Token {
			response := res.SendJson{
				Code: http.StatusUnauthorized,
				Msg:  "Token Error",
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Next()
	})

	// 注册自定义其他路由
	InitMod()
	// 404路由
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404", nil)
	})

	return r
}

func InitMod() {
	user.InitUserRoutes() // 加载用户路由模块
	router.InitRouter(r)  // 将所有方法 注册到路由中
}
