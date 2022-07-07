package model

import (
    . "active/common"
)

/**
"id": "int", //  
"user_id": "string", // 用户id 
"vote_id": "int", // 投票编号 
"vote_option_id": "int", // 投票选项编号 
"value": "string", // 答案 
"vote_time": "string", // 投票时间 
"create_time": "string", //  
"update_time": "string", //  
 */

type VoteUser struct {
    Id int `db:"id" json:"id"`
    UserId string `db:"user_id" json:"user_id"`
    VoteId int `db:"vote_id" json:"vote_id"`
    VoteOptionId int `db:"vote_option_id" json:"vote_option_id"`
    Value string `db:"value" json:"value"`
    VoteTime string `db:"vote_time" json:"vote_time"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewVoteUserModel() *VoteUser {
	return &VoteUser{}
}

func (m *VoteUser) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *VoteUser) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *VoteUser) Update(arg *VoteUser) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *VoteUser) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *VoteUser) SetId(arg int) *VoteUser {
	m.Id = arg
	return m
}

func (m *VoteUser) SetUserId(arg string) *VoteUser {
	m.UserId = arg
	return m
}

func (m *VoteUser) SetVoteId(arg int) *VoteUser {
	m.VoteId = arg
	return m
}

func (m *VoteUser) SetVoteOptionId(arg int) *VoteUser {
	m.VoteOptionId = arg
	return m
}

func (m *VoteUser) SetValue(arg string) *VoteUser {
	m.Value = arg
	return m
}

func (m *VoteUser) SetVoteTime(arg string) *VoteUser {
	m.VoteTime = arg
	return m
}

func (m *VoteUser) SetCreateTime(arg string) *VoteUser {
	m.CreateTime = arg
	return m
}

func (m *VoteUser) SetUpdateTime(arg string) *VoteUser {
	m.UpdateTime = arg
	return m
}

func (m VoteUser) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "user_id": m.UserId, 
        "vote_id": m.VoteId, 
        "vote_option_id": m.VoteOptionId, 
        "value": m.Value, 
        "vote_time": m.VoteTime, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m VoteUser) Translates() map[string]string {
	return map[string]string{ 
        "id": "", 
        "user_id": "用户id", 
        "vote_id": "投票编号", 
        "vote_option_id": "投票选项编号", 
        "value": "答案", 
        "vote_time": "投票时间", 
        "create_time": "", 
        "update_time": "", 
	}
}