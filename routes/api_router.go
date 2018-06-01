package routes

import (
	"github.com/gin-gonic/gin"
	"wechat/controllers"
	"wechat/filters/auth"
)

func RegisterApiRouter(router *gin.Engine) {

	//api := router.Group("/api")
	//api.GET("/index", controllers.IndexApi)

	// 微信sdk

	// api
	// - 管理公众号信息
	// - 推送消息到指定公众号/小程序的指定用户 队列化
	// - 管理 access_token
	// - 微信登录
	// - 获取微信小程序码
	// - 微信支付

	// TODO:
	// - 中间件: 指定内部ip访问
	// - 不操作业务数据库，业务隔离
	// - 多公众号管理，公众号配置化
	// - 管理后台界面

	// 请求参数：
	// - 指定公众号id

	// 特性
	// - 支持http, rpc
	// - 分布式etcd
	// - go-client 客户端支持
	// - c10k并发

	api := router.Group("/api")
	api.POST("/call", controllers.CallMethod)
}
