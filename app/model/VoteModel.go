package model

import (
    . "active/common"
)

/**
"id": "int", //  
"name": "string", // 投票名称 
"code": "string", // 投票标识 
"introduction": "string", // 投票介绍 
"rule": "string", // 投票规则 
"share_times": "int", // 分享次数 
"join_times": "int", // 参与次数 
"visits_times": "int", // 访问次数 
"min_select": "int", // 最少选择数量 
"max_select": "int", // 最大选择数量 
"day_vote_limit": "int", // 每日限制投票次数 
"start_time": "string", // 开始时间 
"end_time": "string", // 结束时间 
"create_time": "string", //  
"update_time": "string", //  
 */

type Vote struct {
    Id int `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
    Code string `db:"code" json:"code"`
    Introduction string `db:"introduction" json:"introduction"`
    Rule string `db:"rule" json:"rule"`
    ShareTimes int `db:"share_times" json:"share_times"`
    JoinTimes int `db:"join_times" json:"join_times"`
    VisitsTimes int `db:"visits_times" json:"visits_times"`
    MinSelect int `db:"min_select" json:"min_select"`
    MaxSelect int `db:"max_select" json:"max_select"`
    DayVoteLimit int `db:"day_vote_limit" json:"day_vote_limit"`
    StartTime string `db:"start_time" json:"start_time"`
    EndTime string `db:"end_time" json:"end_time"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewVoteModel() *Vote {
	return &Vote{}
}

func (m *Vote) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Vote) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Vote) Update(arg *Vote) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Vote) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *Vote) SetId(arg int) *Vote {
	m.Id = arg
	return m
}

func (m *Vote) SetName(arg string) *Vote {
	m.Name = arg
	return m
}

func (m *Vote) SetCode(arg string) *Vote {
	m.Code = arg
	return m
}

func (m *Vote) SetIntroduction(arg string) *Vote {
	m.Introduction = arg
	return m
}

func (m *Vote) SetRule(arg string) *Vote {
	m.Rule = arg
	return m
}

func (m *Vote) SetShareTimes(arg int) *Vote {
	m.ShareTimes = arg
	return m
}

func (m *Vote) SetJoinTimes(arg int) *Vote {
	m.JoinTimes = arg
	return m
}

func (m *Vote) SetVisitsTimes(arg int) *Vote {
	m.VisitsTimes = arg
	return m
}

func (m *Vote) SetMinSelect(arg int) *Vote {
	m.MinSelect = arg
	return m
}

func (m *Vote) SetMaxSelect(arg int) *Vote {
	m.MaxSelect = arg
	return m
}

func (m *Vote) SetDayVoteLimit(arg int) *Vote {
	m.DayVoteLimit = arg
	return m
}

func (m *Vote) SetStartTime(arg string) *Vote {
	m.StartTime = arg
	return m
}

func (m *Vote) SetEndTime(arg string) *Vote {
	m.EndTime = arg
	return m
}

func (m *Vote) SetCreateTime(arg string) *Vote {
	m.CreateTime = arg
	return m
}

func (m *Vote) SetUpdateTime(arg string) *Vote {
	m.UpdateTime = arg
	return m
}

func (m Vote) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "name": m.Name, 
        "code": m.Code, 
        "introduction": m.Introduction, 
        "rule": m.Rule, 
        "share_times": m.ShareTimes, 
        "join_times": m.JoinTimes, 
        "visits_times": m.VisitsTimes, 
        "min_select": m.MinSelect, 
        "max_select": m.MaxSelect, 
        "day_vote_limit": m.DayVoteLimit, 
        "start_time": m.StartTime, 
        "end_time": m.EndTime, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m Vote) Translates() map[string]string {
	return map[string]string{ 
        "id": "", 
        "name": "投票名称", 
        "code": "投票标识", 
        "introduction": "投票介绍", 
        "rule": "投票规则", 
        "share_times": "分享次数", 
        "join_times": "参与次数", 
        "visits_times": "访问次数", 
        "min_select": "最少选择数量", 
        "max_select": "最大选择数量", 
        "day_vote_limit": "每日限制投票次数", 
        "start_time": "开始时间", 
        "end_time": "结束时间", 
        "create_time": "", 
        "update_time": "", 
	}
}