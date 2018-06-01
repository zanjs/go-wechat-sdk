package caller

import "wechat/module/wechat"

var Call = map[string]func(accountid int, data map[string]string){
	"GetNewAccessToken" : wechat.GetNewAccessToken,
	"GetWebOauthAccessToken" : wechat.GetWebOauthAccessToken,
	"RefreshWebOauthAccessToken" : wechat.RefreshWebOauthAccessToken,
	"GetWebOauthUserinfo" : wechat.GetWebOauthUserinfo,
	"CheckWebOauthAccessTokenEffective" : wechat.CheckWebOauthAccessTokenEffective,
	"SendTemplateMessage" : wechat.SendTemplateMessage,
	"WxappOauth" : wechat.WxappOauth,
	"GetWxappCode" : wechat.GetWxappCode,
	"GetWxappCodeUnlimit" : wechat.GetWxappCodeUnlimit,
	"GetWxappCodeQrcode" : wechat.GetWxappCodeQrcode,
	"SendWxappTemplateMessage" : wechat.SendWxappTemplateMessage,
	"PayUnifiedOrder" : wechat.PayUnifiedOrder,
}