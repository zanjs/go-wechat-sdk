package controllers

import (
	// "github.com/gin-gonic/gin"
	"wechat/module/wechat"
)

func GetNewAccessToken()  {
	wechat.GetNewAccessToken()
	return
}

func GetWebOauthAccessToken()  {
	wechat.GetWebOauthAccessToken()
	return
}

func RefreshWebOauthAccessToken()  {
	wechat.RefreshWebOauthAccessToken()
	return
}

func GetWebOauthUserinfo()  {
	wechat.GetWebOauthUserinfo()
	return
}

func CheckWebOauthAccessTokenEffective()  {
	wechat.CheckWebOauthAccessTokenEffective()
	return
}

func SendTemplateMessage()  {
	wechat.SendTemplateMessage()
	return
}

func WxappOauth()  {
	wechat.WxappOauth()
	return
}

func GetWxappCode()  {
	wechat.GetWxappCode()
	return
}

func GetWxappCodeUnlimit()  {
	wechat.GetWxappCodeUnlimit()
	return
}

func GetWxappCodeQrcode()  {
	wechat.GetWxappCodeQrcode()
	return
}

func SendWxappTemplateMessage()  {
	wechat.SendWxappTemplateMessage()
	return
}

func PayUnifiedOrder()  {
	wechat.PayUnifiedOrder()
	return
}