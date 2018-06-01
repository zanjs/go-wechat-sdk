package controllers

import (
	 "github.com/gin-gonic/gin"
	"wechat/module/caller"
	"strconv"
)

func CallMethod(c *gin.Context)  {
	accountId, err := strconv.Atoi(c.PostForm("accountId"))
	if err != nil {
		return
	}
	method := c.PostForm("method")
	if method == "" {
		return
	}
	caller.Call(c, accountId)
	return
}