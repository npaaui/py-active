package wxhelper

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/widuu/gojson"

	. "active/common"
)

const (
	redirectOauthURL       = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	webAppRedirectOauthURL = "https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	webAccessTokenURL      = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL  = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoURL            = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	userInfoWebURL         = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	checkAccessTokenURL    = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	accessTokenURL         = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	getTicketURL           = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	urlSendTemplateMsg     = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s" //发送模版消息地址
	MenuURL                = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	MenuDelURL             = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"         //删除菜单
	customMsgURL           = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s" //客服消息接口
	ShortUrl               = "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"
	listMsg                = "https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=%s"
	openList               = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s"
)

// 普通access_token
type ResAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// 获取普通access_token
//测试公众号测试   √
func GetAccessToken() (res ResAccessToken, err error) {
	urlStr := fmt.Sprintf(accessTokenURL, AppId, AppSecret)
	fmt.Println("urlStr:", urlStr)
	body, errcode := HttpGet(urlStr)
	if errcode != 200 {
		fmt.Println("GetUserAccessToken HttpGet errcode:", errcode)
		return
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println("GetUserAccessToken Unmarshal err:", err)
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return
}

//微信用户授权 获取授权url  √
// 获取redirectOauthURL链接。其在微信中跳转后可以获取code
// 因为微信转码会造成部分链接参数丢失的情况，使用urlEncode对链接进行处理
func RedirectOauthUrl(redirectUrl string) string {
	// url encode
	v := url.Values{}
	v.Add("redirectUrl", redirectUrl) // 添加map
	encodeUrl := v.Encode()
	encodeUrl = strings.TrimLeft(encodeUrl, "redirectUrl=") //去掉url中多余的字符串
	urlStr := fmt.Sprintf(redirectOauthURL, AppId, encodeUrl, "snsapi_userinfo", "123")
	return urlStr
}

//三方公众号授权 https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func RedirectOauthWebUrl(redirectUrl string) string {
	v := url.Values{}
	v.Add("redirectUrl", redirectUrl) // 添加map
	encodeUrl := v.Encode()
	encodeUrl = strings.TrimLeft(encodeUrl, "redirectUrl=") //去掉url中多余的字符串
	urlStr := fmt.Sprintf(webAppRedirectOauthURL, AppId, encodeUrl, "snsapi_login", "321")
	return urlStr
}

type WxUserInfo struct {
	ID         int        `json:"id"`
	Openid     string     `json:"openid"`
	Nickname   string     `json:"nickname"`
	Headimgurl string     `json:"headimgurl"`
	Sex        int        `json:"sex"`
	Province   string     `json:"province"`
	City       string     `json:"city"`
	Country    string     `json:"country"`
	Name       string     `json:"name"`
	Mobile     string     `json:"mobile"`
	Address    string     `json:"address"`
	Subscribe  int        `json:"subscribe"`
	Unionid    string     `json:"unionid"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Errcode    int        `json:"errcode"`
	Errmsg     string     `json:"errmsg"`
}

// 调用微信接口获取用户信息  √
func GetWxUserInfo(openID string) (res WxUserInfo, err error) {
	accessToken := getAccessToken()
	if accessToken == "" {
		return
	}
	urlStr := fmt.Sprintf(userInfoURL, accessToken, openID)
	body, errcode := HttpGet(urlStr)
	fmt.Printf("getWxUserInfo: %+v; \n", body)
	if errcode != 200 {
		return
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return
}

// 网页授权access_token
type ResWebAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Openid       string `json:"openid"`
	Unionid      string `json:"unionid"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

//通过code换取网页授权access_token
func getWebAccessToken(code string) (res ResWebAccessToken, err error) {
	//缓存
	urlStr := fmt.Sprintf(webAccessTokenURL, AppId, AppSecret, code)
	body, errcode := HttpGet(urlStr)
	if errcode != 200 {
		return
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetWebAccessToken error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return
}

func GeWxUserInfoByWebCode(code string) (res WxUserInfo, openId string, err error) {
	accessToken, err := getWebAccessToken(code)
	if err != nil || accessToken.AccessToken == "" {
		err = fmt.Errorf("getWebAccessToken err: %+v", err)
		return
	}
	urlStr := fmt.Sprintf(userInfoWebURL, accessToken.AccessToken, accessToken.Openid)
	body, errcode := HttpGet(urlStr)
	fmt.Printf("geWxUserInfoByWebCode: %+v; \n", body)
	if errcode != 200 {
		err = fmt.Errorf("GeWxUserInfoByWebCode errcode: %v", errcode)
		return
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		err = fmt.Errorf("WxUserInfo Unmarshal err: %v", err)
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return res, accessToken.Openid, err
}

type ResTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}

//申请授权页
//开票等页面
func getTicket() (ticket string) {
	var res ResTicket
	ret, err := CacheServiceInst.Get(CACHE_WX_JS_TICKET)
	if err != nil {
		accessToken := getAccessToken()
		if accessToken == "" {
			err = fmt.Errorf("GetTicket getAccessToken err")
			return
		}
		urlStr := fmt.Sprintf(getTicketURL, accessToken)
		body, errcode := HttpGet(urlStr)
		if errcode != 200 {
			fmt.Println("GetTicket  HttpGet err:", err)
			return
		}
		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			fmt.Println("GetTicket Unmarshal err:", err)
			return
		}
		if res.Errcode != 0 {
			err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", res.Errcode, res.Errmsg)
			return
		}
		//添加缓存 时间为3600s
		CacheServiceInst.Save(CACHE_WX_JS_TICKET, res.Ticket, 3600)
		return res.Ticket
	} else {
		return ret
	}
	return
}

//签名，用于wap分享
func GetWxConfig(url string) map[string]interface{} {
	ticket := getTicket()
	ti := time.Now().Unix()
	timestamp := fmt.Sprintf("%d", ti)
	wxnonceStr := RandomStr(16)
	//这里参数的顺序要按照key值ASCII码升序排序
	sha1Str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, wxnonceStr, timestamp, url)
	signature := getSha1(sha1Str)
	//fmt.Println("signature:", signature)
	signPackage := map[string]interface{}{
		"app_id":     AppId,
		"nonce_str":  wxnonceStr,
		"timestamp":  ti,
		"signature":  signature,
		"raw_string": sha1Str,
	}
	return signPackage
}

func getSha1(str string) (sha1Str string) {
	h := sha1.New()
	io.WriteString(h, str)
	//fmt.Println("sha1Str:", str)
	sha1Str = fmt.Sprintf("%x", h.Sum(nil))
	return
}

type ResCreateMenu struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//生成菜单
func CreateMenu(meunList string) bool {
	accessToken := getAccessToken()
	if accessToken == "" {
		fmt.Println("getWebAccessToken err")
		return false
	}
	urlStr := fmt.Sprintf(MenuURL, accessToken)
	ret, err := HttpRequest(HttpReq{
		Url:    urlStr,
		IsPost: true,
		IsJson: true,
		Data:   []byte(meunList),
	})
	if err != nil {
		fmt.Println("CreateMenu err:", err)
		return false
	}
	var res ResCreateMenu
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		fmt.Println("ResCreateMenu Unmarshal err:", err)
		return false
	}
	if res.Errcode != 0 {
		return false
	}
	return true
}

func DelMenu() bool {
	//MenuDelURL
	accessToken := getAccessToken()
	if accessToken == "" {
		fmt.Println("getWebAccessToken err")
		return false
	}
	urlStr := fmt.Sprintf(MenuDelURL, accessToken)
	ret, errcode := HttpGet(urlStr)
	if errcode != 200 {
		fmt.Println("DelMenu err:", errcode)
		return false
	}
	var res ResCreateMenu
	err := json.Unmarshal([]byte(ret), &res)
	if err != nil {
		fmt.Println("ResCreateMenu Unmarshal err:", err)
		return false
	}
	if res.Errcode != 0 {
		return false
	}
	return true
}

type ResShortUrl struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	ShortUrl string `json:"short_url"`
}

func MakeShortUrl(longUrl string) (short string) {
	accessToken := getAccessToken()
	content := map[string]interface{}{
		"action":   "long2short",
		"long_url": longUrl,
	}
	data, _ := json.Marshal(content)
	urlStr := fmt.Sprintf(ShortUrl, accessToken)
	ret, err := HttpRequest(HttpReq{
		Url:    urlStr,
		IsPost: true,
		IsJson: true,
		Data:   data,
	})
	if err != nil {
		fmt.Println("MakeShortUrl err:", err)
		return
	}
	var res ResShortUrl
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		fmt.Println("Res Unmarshal err:", err)
		return
	}
	if res.Errcode != 0 {
		return
	}
	short = res.ShortUrl
	return
}

type SendCustomMsgParam struct {
	OpenId  string ` json:"open_id"  `
	Content string
}

func SendCustomMsg(param *SendCustomMsgParam) (errInt int) {
	//获取access_token
	accessToken := getAccessToken()
	if accessToken == "" {
		return -1
	}
	urlPath := fmt.Sprintf(customMsgURL, accessToken)
	textCont := map[string]interface{}{
		"content": param.Content,
	}
	//格式化消息体
	paramData := map[string]interface{}{
		"touser":  param.OpenId,
		"msgtype": "text",
		"text":    textCont,
	}
	contentJson, _ := json.Marshal(paramData)
	//消息发送
	sendFlag := 0
	//fmt.Println("发送微信客服消息:", string(contentJson))
	ret, err := HttpRequest(HttpReq{
		Url:    urlPath,
		IsPost: true,
		IsJson: true,
		Data:   contentJson,
	})
	if err != nil {
		fmt.Println("SendCustomMsg err:", err)
		return -1
	}
	fmt.Println("发送微信客服消息结果:", string(ret))

	errData := gojson.Json(string(ret)).Getdata()
	errcode, ok := errData["errcode"].(float64)
	if ok {
		errcodeInt := int(errcode)
		if errcodeInt != 0 {
			fmt.Println("发送微信客服消息失败 err:", errcodeInt, errData["errmsg"].(string))
			return errcodeInt
		} else {
			sendFlag = 1
		}
	} else {
		return -1
	}
	return sendFlag
}

type SendCustomNewsParam struct {
	OpenId  string ` json:"open_id"  `
	Title   string
	Content string
	Url     string
	Img     string
}

//发送图文消息（点击跳转到外链） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008
func SendCustomNews(param *SendCustomNewsParam) (errInt int) {
	//获取access_token
	accessToken := getAccessToken()
	if accessToken == "" {
		return -1
	}
	articles := []map[string]interface{}{}
	urlPath := fmt.Sprintf(customMsgURL, accessToken)
	textCont := map[string]interface{}{
		"title":       param.Title,
		"description": param.Content,
		"url":         param.Url,
		"picurl":      param.Img,
	}
	articles = append(articles, textCont)
	//格式化消息体
	paramData := map[string]interface{}{
		"touser":  param.OpenId,
		"msgtype": "news",
		"news": map[string]interface{}{
			"articles": articles,
		},
	}
	contentJson, _ := json.Marshal(paramData)
	//消息发送
	sendFlag := 0
	fmt.Println("发送微信客服消息:", string(contentJson))
	ret, err := HttpRequest(HttpReq{
		Url:    urlPath,
		IsPost: true,
		IsJson: true,
		Data:   contentJson,
	})
	if err != nil {
		fmt.Println("SendCustomNews err:", err)
		return -1
	}
	fmt.Println("发送微信客服消息结果:", string(ret))

	errData := gojson.Json(string(ret)).Getdata()
	errcode, ok := errData["errcode"].(float64)
	if ok {
		errcodeInt := int(errcode)
		if errcodeInt != 0 {
			fmt.Println("发送微信客服消息失败 err:", errcodeInt, errData["errmsg"].(string))
			return errcodeInt
		} else {
			sendFlag = 1
		}
	} else {
		return -1
	}
	return sendFlag
}

////////////////////////////////////////对外接口///////////////////////////////////////////////

//获取access_token 引入缓存 √
func getAccessToken() string {
	var accessToken ResAccessToken
	ret, err := CacheServiceInst.Get(CACHE_WX_ACCESS_TOKEN + AppId)
	if err != nil || ret == "" {
		accessToken, err = GetAccessToken()
		if err != nil {
			fmt.Println("GetAccessToken err：", err)
			return ""
		}
		token := accessToken.AccessToken
		//添加缓存 时间为3600s
		CacheServiceInst.Save(CACHE_WX_ACCESS_TOKEN+AppId, token, 3600)
		return token
	} else {
		return ret
	}
}

type SendTemplateMsgParam struct {
	OpenId      string      ` json:"open_id"  `
	TemplateId  string      `form:"template_id" json:"template_id"  `
	Content     interface{} `form:"content" json:"content"  `
	Url         string
	Miniprogram string
}

//实际发送操作
func SendTemplateMsg(param *SendTemplateMsgParam) (errInt int) {
	//获取access_token
	accessToken := getAccessToken()
	if accessToken == "" {
		return -1
	}
	urlPath := fmt.Sprintf(urlSendTemplateMsg, accessToken)
	//格式化消息体
	paramData := map[string]interface{}{
		"touser":      param.OpenId,
		"template_id": param.TemplateId,
		"data":        param.Content,
	}
	if param.Url != "" {
		paramData["url"] = param.Url
	}
	if param.Miniprogram != "" {
		miniproData := map[string]string{}
		miniproMap := gojson.Json(param.Miniprogram).Getdata()
		AppId, ok := miniproMap["AppId"].(string)
		if !ok {
			return -1
		}
		miniproData["AppId"] = AppId
		pagepath, ok := miniproMap["pagepath"].(string)
		if ok {
			miniproData["pagepath"] = pagepath
		}
		paramData["miniprogram"] = miniproData
	}
	contentJson, _ := json.Marshal(paramData)
	//消息发送
	sendFlag := 0
	//fmt.Println(string(contentJson))
	ret, err := HttpRequest(HttpReq{
		Url:    urlPath,
		IsPost: true,
		IsJson: true,
		Data:   contentJson,
	})
	if err != nil {
		fmt.Println("SendTemplateMsg err:", err)
		return -1
	}
	errData := gojson.Json(string(ret)).Getdata()
	errcode, ok := errData["errcode"].(float64)
	if ok {
		errcodeInt := int(errcode)
		if errcodeInt != 0 {
			fmt.Println("发送微信消息失败 err:", errcodeInt, errData["errmsg"].(string))
			return errcodeInt
		} else {
			sendFlag = 1
		}
	} else {
		return -1
	}
	return sendFlag
}

/////////////////////////////////////////////////////////////
type WxOpenids struct {
	Total   int    `json:"total"`
	Count   int    `json:"count"`
	Data    OpenL  `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type OpenL struct {
	Openid []string `json:"openid"`
}

func GetOpenList() (res WxOpenids, err error) {
	accessToken := getAccessToken()
	urlStr := fmt.Sprintf(openList, accessToken)
	body, errcode := HttpGet(urlStr)
	if errcode != 200 {
		return
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return
	}
	fmt.Println(res)
	return
}

//群发消息
func SendMsgList(content string, openList []string) (errInt int) {
	accessToken := getAccessToken()
	urlPath := fmt.Sprintf(listMsg, accessToken)
	textCont := map[string]interface{}{
		"content": content,
	}
	fmt.Println("urlPath:", urlPath)
	if len(openList) == 0 {
		openid, _ := GetOpenList()
		fmt.Println(openid.Data.Openid)
		openList = openid.Data.Openid
	}
	//格式化消息体
	paramData := map[string]interface{}{
		"touser":  openList,
		"msgtype": "text",
		"text":    textCont,
	}
	contentJson, _ := json.Marshal(paramData)
	//消息发送
	sendFlag := 0
	fmt.Println(string(contentJson))
	ret, err := HttpRequest(HttpReq{
		Url:    urlPath,
		IsPost: true,
		IsJson: true,
		Data:   contentJson,
	})
	if err != nil {
		fmt.Println("SendCustomMsg err:", err)
		return -1
	}
	errData := gojson.Json(string(ret)).Getdata()
	errcode, ok := errData["errcode"].(float64)
	if ok {
		errcodeInt := int(errcode)
		if errcodeInt != 0 {
			fmt.Println("群发消息失败 err:", errcodeInt, errData["errmsg"].(string))
			return errcodeInt
		} else {
			sendFlag = 1
		}
	} else {
		return -1
	}
	return sendFlag
}

//全列表群发消息
func SendMsgListMakeList(content string) (errInt int) {
	openid, _ := GetOpenList()
	openList := openid.Data.Openid
	fmt.Println(openList)
	sendFlag := SendMsgList(content, openList)
	return sendFlag
}
