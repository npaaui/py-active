package dao

import (
	"active/app/model"
	. "active/common"
)

func InfoUserByOpenId(openId string) *model.User {
	user := &model.User{
		OpenId: openId,
	}
	_, err := GetDbEngineIns().MustCols("open_id").Get(user)
	if err != nil {
		panic(NewDbErr(err))
	}
	return user
}