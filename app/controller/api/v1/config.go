package v1

import (
	"active/app/service"
	. "active/common"
	"github.com/gin-gonic/gin"
)

type ConfigCtr struct {
	Srv *service.ConfigSrv
}

func NewConfigCtr() *ConfigCtr {
	return &ConfigCtr{
		Srv: service.NewConfigSrv(),
	}
}

func (t *ConfigCtr) InfoConfig(c *gin.Context) {
	args := &struct {
		Type  string `json:"type"`
		Url string `json:"url"`
	}{}
	_ = ValidateQuery(c, map[string]string{
		"type": "string",
		"url": "string|required",
	}, args)

	data := map[string]interface{}{}
	switch args.Type {
	case "wx_share":
		data = t.Srv.InfoWxShareConfig(args.Url)
		data["title"] = "大奖来了！“寻味饶州”等你投票"
		data["content"] = "“寻味饶州”文和里杯首届旅游美食评选大赛等你来投票！"
		data["img"] = "https://npaaui.oss-cn-hangzhou.aliyuncs.com/py/head.jpeg"
		data["link"] = args.Url
		break
	default:
		break
	}

	ReturnData(c, data)
}
