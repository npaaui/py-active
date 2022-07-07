package model

import (
    . "active/common"
)

/**
"id": "int", // 奖品编号 
"lottery_id": "int", // 活动编号 
"name": "string", // 奖品名称 
"price": "string", // 奖品价格 
"total_num": "int", // 奖品数量 
"winner_num": "int", // 已中奖数量 
"reserve_num": "int", // 备用奖品数量 
"reserve_winner_num": "int", // 备用已中奖数量 
"type": "string", // 奖品类型 ticket 券 goods 物品 money 现金 none 未中奖 score积分 card 话费 
"display_order": "int", // 排序 
"user_limit": "int", // 单个用户限制中奖次数 0为不限制 
"img": "string", // 奖品图片 
"opt": "string", // 可选参数 
"use_desc": "string", // 使用规则 
"create_time": "string", //  
"update_time": "string", //  
 */

type LotteryAward struct {
    Id int `db:"id" json:"id"`
    LotteryId int `db:"lottery_id" json:"lottery_id"`
    Name string `db:"name" json:"name"`
    Price string `db:"price" json:"price"`
    TotalNum int `db:"total_num" json:"total_num"`
    WinnerNum int `db:"winner_num" json:"winner_num"`
    ReserveNum int `db:"reserve_num" json:"reserve_num"`
    ReserveWinnerNum int `db:"reserve_winner_num" json:"reserve_winner_num"`
    Type string `db:"type" json:"type"`
    DisplayOrder int `db:"display_order" json:"display_order"`
    UserLimit int `db:"user_limit" json:"user_limit"`
    Img string `db:"img" json:"img"`
    Opt string `db:"opt" json:"opt"`
    UseDesc string `db:"use_desc" json:"use_desc"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewLotteryAwardModel() *LotteryAward {
	return &LotteryAward{}
}

func (m *LotteryAward) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *LotteryAward) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *LotteryAward) Update(arg *LotteryAward) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *LotteryAward) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *LotteryAward) SetId(arg int) *LotteryAward {
	m.Id = arg
	return m
}

func (m *LotteryAward) SetLotteryId(arg int) *LotteryAward {
	m.LotteryId = arg
	return m
}

func (m *LotteryAward) SetName(arg string) *LotteryAward {
	m.Name = arg
	return m
}

func (m *LotteryAward) SetPrice(arg string) *LotteryAward {
	m.Price = arg
	return m
}

func (m *LotteryAward) SetTotalNum(arg int) *LotteryAward {
	m.TotalNum = arg
	return m
}

func (m *LotteryAward) SetWinnerNum(arg int) *LotteryAward {
	m.WinnerNum = arg
	return m
}

func (m *LotteryAward) SetReserveNum(arg int) *LotteryAward {
	m.ReserveNum = arg
	return m
}

func (m *LotteryAward) SetReserveWinnerNum(arg int) *LotteryAward {
	m.ReserveWinnerNum = arg
	return m
}

func (m *LotteryAward) SetType(arg string) *LotteryAward {
	m.Type = arg
	return m
}

func (m *LotteryAward) SetDisplayOrder(arg int) *LotteryAward {
	m.DisplayOrder = arg
	return m
}

func (m *LotteryAward) SetUserLimit(arg int) *LotteryAward {
	m.UserLimit = arg
	return m
}

func (m *LotteryAward) SetImg(arg string) *LotteryAward {
	m.Img = arg
	return m
}

func (m *LotteryAward) SetOpt(arg string) *LotteryAward {
	m.Opt = arg
	return m
}

func (m *LotteryAward) SetUseDesc(arg string) *LotteryAward {
	m.UseDesc = arg
	return m
}

func (m *LotteryAward) SetCreateTime(arg string) *LotteryAward {
	m.CreateTime = arg
	return m
}

func (m *LotteryAward) SetUpdateTime(arg string) *LotteryAward {
	m.UpdateTime = arg
	return m
}

func (m LotteryAward) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "lottery_id": m.LotteryId, 
        "name": m.Name, 
        "price": m.Price, 
        "total_num": m.TotalNum, 
        "winner_num": m.WinnerNum, 
        "reserve_num": m.ReserveNum, 
        "reserve_winner_num": m.ReserveWinnerNum, 
        "type": m.Type, 
        "display_order": m.DisplayOrder, 
        "user_limit": m.UserLimit, 
        "img": m.Img, 
        "opt": m.Opt, 
        "use_desc": m.UseDesc, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m LotteryAward) Translates() map[string]string {
	return map[string]string{ 
        "id": "奖品编号", 
        "lottery_id": "活动编号", 
        "name": "奖品名称", 
        "price": "奖品价格", 
        "total_num": "奖品数量", 
        "winner_num": "已中奖数量", 
        "reserve_num": "备用奖品数量", 
        "reserve_winner_num": "备用已中奖数量", 
        "type": "奖品类型 ticket 券 goods 物品 money 现金 none 未中奖 score积分 card 话费", 
        "display_order": "排序", 
        "user_limit": "单个用户限制中奖次数 0为不限制", 
        "img": "奖品图片", 
        "opt": "可选参数", 
        "use_desc": "使用规则", 
        "create_time": "", 
        "update_time": "", 
	}
}