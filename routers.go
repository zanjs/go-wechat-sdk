package main

import (
	"wechat/config"
	"github.com/gin-gonic/gin"
	"wechat/filters/auth"
	"wechat/filters"
	routeRegister "wechat/routes"
	"net/http"
)

func initRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TEMPLATE_PATH + "/*") // html模板

	router.Use(gin.Logger())

	router.Use(handleErrors()) // 错误处理
	router.Use(filters.RegisterSession())  // 全局session
	router.Use(filters.RegisterCache())    // 全局cache

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth")) // 全局auth cookie
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))    // 全局auth jwt

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
		return
	})

	routeRegister.RegisterApiRouter(router)

	return router
}
