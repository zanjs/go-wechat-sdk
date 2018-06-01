package caller

import (
	"wechat/module/wechat"
	"github.com/gin-gonic/gin"
	"errors"
	"wechat/module/account"
	"fmt"
)

func Call(c *gin.Context, accountId int) (map[string]string, error){

	method := c.PostForm("method")

	// 数据验证
	if !Validate(method) {
		return nil, errors.New("错误的参数")
	}

	accountInfo := account.GetAccountInfo(accountId)

	var (
		err error
		resData map[string]string
	)

	// 执行结果
	switch c.PostForm("method") {
	case "GetNewAccessToken":
		resData, err = wechat.GetNewAccessToken(accountInfo["appId"], accountInfo["appSecret"])
	case "GetWebOauthAccessToken":
		resData, err = wechat.GetWebOauthAccessToken(accountInfo["appId"], accountInfo["appSecret"], c.PostForm("code"))
	case "RefreshWebOauthAccessToken":
		resData, err = wechat.RefreshWebOauthAccessToken(accountInfo["appId"], c.PostForm("refreshToken"))
	case "GetWebOauthUserinfo":
		resData, err = wechat.GetWebOauthUserinfo(accountInfo["appId"],c.PostForm("lang"),c.PostForm("refreshToken"))
	case "CheckWebOauthAccessTokenValid":
		resData, err = wechat.CheckWebOauthAccessTokenValid(c.PostForm("openId"),c.PostForm("refreshToken"))
	case "SendTemplateMessage":
		resData, err = wechat.SendTemplateMessage(1, map[string]string{})
	case "WxappOauth":
		resData, err = wechat.WxappOauth(accountInfo["appId"], accountInfo["appSecret"], c.PostForm("jsCode"))
	case "GetWxappCode":
		resData, err = wechat.GetWxappCode(map[string]string{
			"path":c.PostForm("path"),
		})
	case "GetWxappCodeUnlimit":
		resData, err = wechat.GetWxappCodeUnlimit(map[string]string{
			"path":c.PostForm("path"),
		})
	case "GetWxappCodeQrcode":
		resData, err = wechat.GetWxappCodeQrcode(map[string]string{
			"path":c.PostForm("path"),
		})
	case "SendWxappTemplateMessage":
		resData, err = wechat.SendWxappTemplateMessage(1, map[string]string{})
	case "PayUnifiedOrder":
		resData, err = wechat.PayUnifiedOrder(1, map[string]string{})
	}

	if err != nil {
		return map[string]string{}, err
	}

	return resData, nil
}

func Validate(method string) bool {
	fmt.Println(method)
	return true
}