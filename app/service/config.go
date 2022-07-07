package service

import (
	. "active/common"
	"active/lib/wxhelper"
)

type ConfigSrv struct{}

func NewConfigSrv() *ConfigSrv {
	return &ConfigSrv{}
}

// 获取微信分享配置
func (s *ConfigSrv) InfoWxShareConfig(url string) (wxShareConf map[string]interface{}) {
	wxhelper.Initialize(Cfg.AppId, Cfg.Secret, Cfg.Token, Cfg.EncodingAESKey)
	wxShareConf = wxhelper.GetWxConfig(url)
	return
}
