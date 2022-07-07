package model

import (
    . "active/common"
)

/**
"id": "int", //  
"vote_id": "int", // 投票编号 
"group": "string", // 分组 
"title": "string", // 选项标题 
"content": "string", // 选项内容 
"img": "string", // 选项图片 
"remark": "string", // 选项备注 
"votes": "int", // 投票数量 
"type": "string", // 选项类型 
"create_time": "string", //  
"update_time": "string", //  
 */

type VoteOption struct {
    Id int `db:"id" json:"id"`
    VoteId int `db:"vote_id" json:"vote_id"`
    Group string `db:"group" json:"group"`
    Title string `db:"title" json:"title"`
    Content string `db:"content" json:"content"`
    Img string `db:"img" json:"img"`
    Remark string `db:"remark" json:"remark"`
    Votes int `db:"votes" json:"votes"`
    Type string `db:"type" json:"type"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewVoteOptionModel() *VoteOption {
	return &VoteOption{}
}

func (m *VoteOption) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *VoteOption) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *VoteOption) Update(arg *VoteOption) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *VoteOption) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *VoteOption) SetId(arg int) *VoteOption {
	m.Id = arg
	return m
}

func (m *VoteOption) SetVoteId(arg int) *VoteOption {
	m.VoteId = arg
	return m
}

func (m *VoteOption) SetGroup(arg string) *VoteOption {
	m.Group = arg
	return m
}

func (m *VoteOption) SetTitle(arg string) *VoteOption {
	m.Title = arg
	return m
}

func (m *VoteOption) SetContent(arg string) *VoteOption {
	m.Content = arg
	return m
}

func (m *VoteOption) SetImg(arg string) *VoteOption {
	m.Img = arg
	return m
}

func (m *VoteOption) SetRemark(arg string) *VoteOption {
	m.Remark = arg
	return m
}

func (m *VoteOption) SetVotes(arg int) *VoteOption {
	m.Votes = arg
	return m
}

func (m *VoteOption) SetType(arg string) *VoteOption {
	m.Type = arg
	return m
}

func (m *VoteOption) SetCreateTime(arg string) *VoteOption {
	m.CreateTime = arg
	return m
}

func (m *VoteOption) SetUpdateTime(arg string) *VoteOption {
	m.UpdateTime = arg
	return m
}

func (m VoteOption) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "vote_id": m.VoteId, 
        "group": m.Group, 
        "title": m.Title, 
        "content": m.Content, 
        "img": m.Img, 
        "remark": m.Remark, 
        "votes": m.Votes, 
        "type": m.Type, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m VoteOption) Translates() map[string]string {
	return map[string]string{ 
        "id": "", 
        "vote_id": "投票编号", 
        "group": "分组", 
        "title": "选项标题", 
        "content": "选项内容", 
        "img": "选项图片", 
        "remark": "选项备注", 
        "votes": "投票数量", 
        "type": "选项类型", 
        "create_time": "", 
        "update_time": "", 
	}
}