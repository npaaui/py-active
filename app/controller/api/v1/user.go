package v1

import (
	"active/app/service"
	. "active/common"
	"github.com/gin-gonic/gin"
)

type UserCtr struct {
	Srv *service.UserSrv
}

func NewUserCtr() *UserCtr {
	return &UserCtr{
		Srv: service.NewUserSrv(),
	}
}

func (t *UserCtr) InsertUser(c *gin.Context) {
	args := &service.InsertUserArgs{}
	ValidatePostJson(c, map[string]string{
		"from": "string|required|enum:wechat",
		"code": "string|required",
	}, args)

	t.Srv.SetContext(c)
	ret, respErr := t.Srv.InsertUserByWx(args)
	if respErr.NotNil() {
		ReturnByRespErr(c, respErr)
		return
	}

	c.SetCookie("openid", ret.OpenId, 86400 * 20, "/", Cfg.Domain, false, true)
	ReturnData(c, ret)
	return
}
