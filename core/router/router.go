package router

import (
	"github.com/gin-gonic/gin"
)

type FnRegisRoute = func(rgPublic *gin.RouterGroup)

type gfnRoutesMod struct {
	FnRegisRoute FnRegisRoute `json:"FnRegisRoute"` // 路由方法
	ModName      string       `json:"ModName"`      // 模块名称
}

// 路由数组
var (
	gfnRoutes []gfnRoutesMod
)

// 注册路由
func ReqisRoute(modName string, fn FnRegisRoute) {
	if fn == nil {
		return
	}
	ModRoute := gfnRoutesMod{
		FnRegisRoute: fn,
		ModName:      modName,
	}
	gfnRoutes = append(gfnRoutes, ModRoute)
}

// 初始化
func InitRouter(router *gin.Engine) *gin.Engine {

	for _, fnRegisRote := range gfnRoutes {
		fnRegisRote.FnRegisRoute(router.Group(fnRegisRote.ModName))
	}

	return router
}
