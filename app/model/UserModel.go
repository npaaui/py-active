package model

import (
    . "active/common"
)

/**
"id": "string", // 用户id 
"mobile": "string", // 电话号码 
"open_id": "string", // openid 
"union_id": "string", // union_id 
"nick_name": "string", // 昵称 
"head_pic": "string", // 头像 
"province": "string", // 省 
"city": "string", // 市 
"sex": "int", // 性别 
"focus": "int", // 是否关注 
"create_time": "string", // 添加时间 
"update_time": "string", // 修改时间 
 */

type User struct {
    Id string `db:"id" json:"id"`
    Mobile string `db:"mobile" json:"mobile"`
    OpenId string `db:"open_id" json:"open_id"`
    UnionId string `db:"union_id" json:"union_id"`
    NickName string `db:"nick_name" json:"nick_name"`
    HeadPic string `db:"head_pic" json:"head_pic"`
    Province string `db:"province" json:"province"`
    City string `db:"city" json:"city"`
    Sex int `db:"sex" json:"sex"`
    Focus int `db:"focus" json:"focus"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewUserModel() *User {
	return &User{}
}

func (m *User) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *User) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *User) Update(arg *User) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *User) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *User) SetId(arg string) *User {
	m.Id = arg
	return m
}

func (m *User) SetMobile(arg string) *User {
	m.Mobile = arg
	return m
}

func (m *User) SetOpenId(arg string) *User {
	m.OpenId = arg
	return m
}

func (m *User) SetUnionId(arg string) *User {
	m.UnionId = arg
	return m
}

func (m *User) SetNickName(arg string) *User {
	m.NickName = arg
	return m
}

func (m *User) SetHeadPic(arg string) *User {
	m.HeadPic = arg
	return m
}

func (m *User) SetProvince(arg string) *User {
	m.Province = arg
	return m
}

func (m *User) SetCity(arg string) *User {
	m.City = arg
	return m
}

func (m *User) SetSex(arg int) *User {
	m.Sex = arg
	return m
}

func (m *User) SetFocus(arg int) *User {
	m.Focus = arg
	return m
}

func (m *User) SetCreateTime(arg string) *User {
	m.CreateTime = arg
	return m
}

func (m *User) SetUpdateTime(arg string) *User {
	m.UpdateTime = arg
	return m
}

func (m User) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "mobile": m.Mobile, 
        "open_id": m.OpenId, 
        "union_id": m.UnionId, 
        "nick_name": m.NickName, 
        "head_pic": m.HeadPic, 
        "province": m.Province, 
        "city": m.City, 
        "sex": m.Sex, 
        "focus": m.Focus, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m User) Translates() map[string]string {
	return map[string]string{ 
        "id": "用户id", 
        "mobile": "电话号码", 
        "open_id": "openid", 
        "union_id": "union_id", 
        "nick_name": "昵称", 
        "head_pic": "头像", 
        "province": "省", 
        "city": "市", 
        "sex": "性别", 
        "focus": "是否关注", 
        "create_time": "添加时间", 
        "update_time": "修改时间", 
	}
}