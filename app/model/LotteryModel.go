package model

import (
    . "active/common"
)

/**
"id": "int", // 编号 
"name": "string", // 名称 
"code": "string", // 标示 
"introduction": "string", // 介绍 
"rules": "string", // 规则 
"remark": "string", // 特别说明 
"share_times": "int", // 分享次数 
"join_times": "int", // 参与次数 
"cost_integral": "int", // 每次消耗积分 
"day_lottery_limit": "int", // 每天抽奖限制 
"user_win_limit": "int", // 用户每天中奖次数限制 
"start_time": "string", // 开始时间 
"end_time": "string", // 结束时间 
"create_time": "string", // 添加时间 
"update_time": "string", // 更新时间 
 */

type Lottery struct {
    Id int `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
    Code string `db:"code" json:"code"`
    Introduction string `db:"introduction" json:"introduction"`
    Rules string `db:"rules" json:"rules"`
    Remark string `db:"remark" json:"remark"`
    ShareTimes int `db:"share_times" json:"share_times"`
    JoinTimes int `db:"join_times" json:"join_times"`
    CostIntegral int `db:"cost_integral" json:"cost_integral"`
    DayLotteryLimit int `db:"day_lottery_limit" json:"day_lottery_limit"`
    UserWinLimit int `db:"user_win_limit" json:"user_win_limit"`
    StartTime string `db:"start_time" json:"start_time"`
    EndTime string `db:"end_time" json:"end_time"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewLotteryModel() *Lottery {
	return &Lottery{}
}

func (m *Lottery) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Lottery) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Lottery) Update(arg *Lottery) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Lottery) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *Lottery) SetId(arg int) *Lottery {
	m.Id = arg
	return m
}

func (m *Lottery) SetName(arg string) *Lottery {
	m.Name = arg
	return m
}

func (m *Lottery) SetCode(arg string) *Lottery {
	m.Code = arg
	return m
}

func (m *Lottery) SetIntroduction(arg string) *Lottery {
	m.Introduction = arg
	return m
}

func (m *Lottery) SetRules(arg string) *Lottery {
	m.Rules = arg
	return m
}

func (m *Lottery) SetRemark(arg string) *Lottery {
	m.Remark = arg
	return m
}

func (m *Lottery) SetShareTimes(arg int) *Lottery {
	m.ShareTimes = arg
	return m
}

func (m *Lottery) SetJoinTimes(arg int) *Lottery {
	m.JoinTimes = arg
	return m
}

func (m *Lottery) SetCostIntegral(arg int) *Lottery {
	m.CostIntegral = arg
	return m
}

func (m *Lottery) SetDayLotteryLimit(arg int) *Lottery {
	m.DayLotteryLimit = arg
	return m
}

func (m *Lottery) SetUserWinLimit(arg int) *Lottery {
	m.UserWinLimit = arg
	return m
}

func (m *Lottery) SetStartTime(arg string) *Lottery {
	m.StartTime = arg
	return m
}

func (m *Lottery) SetEndTime(arg string) *Lottery {
	m.EndTime = arg
	return m
}

func (m *Lottery) SetCreateTime(arg string) *Lottery {
	m.CreateTime = arg
	return m
}

func (m *Lottery) SetUpdateTime(arg string) *Lottery {
	m.UpdateTime = arg
	return m
}

func (m Lottery) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "name": m.Name, 
        "code": m.Code, 
        "introduction": m.Introduction, 
        "rules": m.Rules, 
        "remark": m.Remark, 
        "share_times": m.ShareTimes, 
        "join_times": m.JoinTimes, 
        "cost_integral": m.CostIntegral, 
        "day_lottery_limit": m.DayLotteryLimit, 
        "user_win_limit": m.UserWinLimit, 
        "start_time": m.StartTime, 
        "end_time": m.EndTime, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m Lottery) Translates() map[string]string {
	return map[string]string{ 
        "id": "编号", 
        "name": "名称", 
        "code": "标示", 
        "introduction": "介绍", 
        "rules": "规则", 
        "remark": "特别说明", 
        "share_times": "分享次数", 
        "join_times": "参与次数", 
        "cost_integral": "每次消耗积分", 
        "day_lottery_limit": "每天抽奖限制", 
        "user_win_limit": "用户每天中奖次数限制", 
        "start_time": "开始时间", 
        "end_time": "结束时间", 
        "create_time": "添加时间", 
        "update_time": "更新时间", 
	}
}