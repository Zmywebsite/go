package res

import (
	"github.com/gin-gonic/gin"
)

// 返回响应
func Send(ctx *gin.Context, ResponseJson SendJson) {
	ctx.AbortWithStatusJSON(ResponseJson.Code, ResponseJson)
}

type SendJson struct {
	Code int         `json:"code"`           // 请求结 (状态码)
	Msg  string      `json:"msg"`            // 请求结果 (中文)
	Data interface{} `json:"data,omitempty"` // 请求结果 (如果是数组, 在对象中添加list字段)
}

