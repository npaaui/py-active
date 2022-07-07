package wxhelper

import (
	"regexp"
	log "github.com/sirupsen/logrus"
	"encoding/base64"
	"os"
)

// TConfig 配置
var (
	AppId          string // 应用ID
	AppSecret      string // 应用密钥
	Token          string // 令牌
	EncodingAESKey []byte // 消息加解密密钥
)

// Initialize 配置并初始化
func Initialize(appId, appSecret, token, encodingAESKey string) {
	//if matched, err := regexp.MatchString("^gh_[0-9a-f]{12}$", originId); err != nil || !matched {
	//	log.Fatalf("originId format error: %s", err)
	//}
	if matched, err := regexp.MatchString("^wx[0-9a-f]{16}$", appId); err != nil || !matched {
		log.Fatalf("appId format error: %s", err)
	}
	if matched, err := regexp.MatchString("^[0-9a-f]{32}$", appSecret); err != nil || !matched {
		log.Fatalf("appSecret format error: %s", err)
	}
	//服务器
	if token != "" {
		if matched, err := regexp.MatchString("^[0-9a-zA-Z]{3,32}$", token); err != nil || !matched {
			log.Fatalf("token format error: %s", err)
		}
		if matched, err := regexp.MatchString("^[0-9a-zA-Z]{43}$", encodingAESKey); err != nil || !matched {
			log.Fatalf("encodingAESKey format error: %s", err)
		}
		var err error
		EncodingAESKey, err = base64.StdEncoding.DecodeString(encodingAESKey + "=")
		if err != nil {
			log.Info("appSecret config error: %s", err)
			os.Exit(1)
		}
	}

	AppId = appId         // 应用ID
	AppSecret = appSecret // 应用密钥
	Token = token         // 令牌

	CacheServiceInst.Init()
	// refresh access token
	//RefreshAccessToken(AppId, AppSecret)
}
