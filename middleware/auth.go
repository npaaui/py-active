package middleware

import (
	"active/app/dao"
	. "active/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		openId, err := g.Cookie("openid")
		if err != nil {
			LogErr(g.GetString("req_no"), "getOpenidFromCookie", LogLevelWarning, fmt.Errorf("cookie获取失败: %v", err))
		}
		Log(g.GetString("req_no"), "getOpenidFromCookie", LogLevelNormal, "open_id: " + openId)
		user := dao.InfoUserByOpenId(openId)
		if user.Id == "" || user.NickName == "" {
			ReturnErrMsg(g, ErrAuth, "授权处理中,请稍后重试")
			g.Abort()
			return
		}
		g.Set("user_id", user.Id)
		g.Next()
	}
}
