package user

import (
	"net/http"
	"vi-server/core/res"
	"vi-server/core/router"

	_ "vi-server/docs"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	router.ReqisRoute("user", func(Router *gin.RouterGroup) {
		// 用户登录
		Router.POST("/login", func(ctx *gin.Context) {
			type User struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			var (
				user User
			)
			ctx.ShouldBindJSON(&user)
			println("用户名", user.Username)
			print("密码", user.Password)
			islogin := userLogin(user.Username, user.Password)
			if islogin {
				res.Send(ctx, res.SendJson{
					Code: http.StatusOK,
					Msg:  "SUCCESS",
				})
			} else {
				res.Send(ctx, res.SendJson{
					Code: http.StatusBadRequest,
					Msg:  "NAME OR PASSWORD ERROR",
				})
			}
		})

		// 添加用户
		Router.GET("/add", func(ctx *gin.Context) {
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")
			if username == "" {
				res.Send(ctx, res.SendJson{
					Code: http.StatusBadRequest,
					Msg:  "USERNAME ERROR",
				})
			}
			if password == "" {
				res.Send(ctx, res.SendJson{
					Code: http.StatusBadRequest,
					Msg:  "PASSWORD ERROR",
				})
			}

			if addUser(username, password) {
				res.Send(ctx, res.SendJson{
					Code: http.StatusOK,
					Msg:  "SUCCESS",
				})
			} else {
				res.Send(ctx, res.SendJson{
					Code: http.StatusInternalServerError,
					Msg:  "SERVER ERROR",
				})
			}
		})
	})
}
