package wechat

import (
	"net/http"
	"io"
	"errors"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"compress/gzip"
)

// 内部api
// - 动态获取公众号数据appid, appsecret, payid, paysecret
// - 构建签名, 签名算法
// - post 请求
// - get 请求
// - 获取access_token

// 文档：

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1445241432  微信公众号
// https://developers.weixin.qq.com/miniprogram/dev/api/qrcode.html  微信小程序
// https://pay.weixin.qq.com/wiki/doc/api/index.html  微信支付


// ---------------------------
// API常数
// ---------------------------

const (
	GET_ACCESS_TOKEN_API = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET" // 获取access_token

	// 网页授权

	GET_WEB_OAUTH_ACCESS_TOKEN = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code" // 获取特殊的网页授权access_token
	REFRESH_WEB_OAUTH_ACCESS_TOKEN = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN" // 刷新token
	GET_WEB_OAUTH_USERINFO = "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN" // 拉取用户信息(需scope为 snsapi_userinfo)
	CHECK_WEB_OAUTH_ACCESS_TOKEN_EFFECTIVE = "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID" // 检验token有效性

	// 模板消息

	SEND_TEMPLATE_MESSAGE = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN" // 发送模板消息

	// 小程序登录

	WXAPP_OAUTH = "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code" // 小程序获取sessionkey
	GET_WXAPP_CODE = "https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN" // 获取小程序码
	GET_WXAPP_CODE_UNLIMIT = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN" // 获取小程序码
	GET_WXAPP_CODE_QRCODE = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN" // 获取小程序二维码
	SEND_WXAPP_TEMPLATE_MESSAGE  = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN" // 发送小程序服务通知

	// 微信支付

	PAY_UNIFIED_ORDER = "https://api.mch.weixin.qq.com/pay/unifiedorder" // 下订单
)

func GetNewAccessToken()  {
	return
}

func GetWebOauthAccessToken()  {
	return
}

func RefreshWebOauthAccessToken()  {
	return
}

func GetWebOauthUserinfo()  {
	return
}

func CheckWebOauthAccessTokenEffective()  {
	return
}

func SendTemplateMessage()  {
	return
}

func WxappOauth()  {
	return
}

func GetWxappCode()  {
	return
}

func GetWxappCodeUnlimit()  {
	return
}

func GetWxappCodeQrcode()  {
	return
}

func SendWxappTemplateMessage()  {
	return
}

func PayUnifiedOrder()  {
	return
}

// ---------------------------
// sdk内部Api
// ---------------------------

func GetToken()  {
	return
}


func MakeGetReq(url string, data map[string]string) (map[string]interface{}, error) {

	var count = 0
	for k, v := range data {
		if count == 0 {
			url += "?" + k + "=" + v
		} else {
			url += "&" + k + "=" + v
		}
		count++
	}

	res, err := http.Get(url)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if res.StatusCode != 200 {
		return map[string]interface{}{}, errors.New("网络错误")
	}

	var reader io.ReadCloser
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return map[string]interface{}{}, err
		}
	} else {
		reader = res.Body
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var resJsonData map[string]interface{}
	err = json.Unmarshal(body, &resJsonData)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return resJsonData, nil
}

func MakePostReq(url string, postData map[string]interface{}, contentType string) (map[string]interface{}, error) {
	jsonData ,jsonErr := json.Marshal(postData)
	if jsonErr != nil {
		return map[string]interface{}{}, jsonErr
	}

	res, _ := http.Post(url, contentType, bytes.NewBuffer(jsonData))

	var (
		reader io.ReadCloser
		err error
	)
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return map[string]interface{}{}, err
		}
	} else {
		reader = res.Body
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var resJsonData map[string]interface{}
	err = json.Unmarshal(body, &resJsonData)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return resJsonData, nil
}