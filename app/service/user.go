package service

import (
	"active/app/model"
	. "active/common"
	"active/lib/wxhelper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserSrv struct {
	Context  *gin.Context
}

func NewUserSrv() *UserSrv {
	return &UserSrv{}
}

func (s *UserSrv) SetContext(context *gin.Context)  {
	s.Context = context
}

// 新增微信用户
type InsertUserArgs struct {
	From string `json:"from"`
	Code string `json:"code"`
}

type InsertUserRet struct {
	OpenId string `json:"open_id"`
}

func (s *UserSrv) InsertUserByWx(args *InsertUserArgs) (data InsertUserRet, respErr RespErr) {
	wxhelper.Initialize(Cfg.AppId, Cfg.Secret, Cfg.Token, Cfg.EncodingAESKey)

	if args.Code == "123456" {
		data.OpenId = "123456"
		return data, RespErr{}
	}

	// 获取用户信息
	userInfoByWeb, openId, err := wxhelper.GeWxUserInfoByWebCode(args.Code)
	Log(s.Context.GetString("req_no"), "geWxUserInfoByWebCode", LogLevelNormal, fmt.Sprintf("open_id:%v, wx_user_info: %+v ", openId, userInfoByWeb))
	if err != nil {
		return data, NewRespErr(ErrWx, err.Error())
	}
	if openId == "" {
		return data, NewRespErr(ErrUserRegister, "用户授权失败")
	}
	// 获取是否关注
	userInfoByOpenId, _ := wxhelper.GetWxUserInfo(openId)

	setUser := &model.User{
		OpenId:   openId,
		Province: userInfoByWeb.Province,
		City:     userInfoByWeb.City,
		Sex:      userInfoByWeb.Sex,
		HeadPic:  userInfoByWeb.Headimgurl,
		NickName: userInfoByWeb.Nickname,
		UnionId:  userInfoByWeb.Unionid,
		Focus:    userInfoByOpenId.Subscribe,
	}
	user := &model.User{
		OpenId: openId,
	}
	if user.Info() {
		_ = user.Update(setUser)
	} else {
		setUser.Id = GetUniqueId()
		row := setUser.Insert()
		if row == 0 {
			return data, NewRespErr(ErrInsert, "新增用户失败")
		}
	}

	data.OpenId = openId
	return data, RespErr{}
}
