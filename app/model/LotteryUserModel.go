package model

import (
    . "active/common"
)

/**
"id": "int", //  
"user_id": "string", // 用户编号 
"mobile": "string", // 手机号码 
"lottery_id": "int", // 活动编号 
"lottery_award_id": "int", // 奖品编号 
"use_gifts_sn": "string", // 兑奖号码 
"cost_integral": "int", // 消耗的积分 
"status": "string", // 发奖状态 INIT 未发奖 DONE 已发奖 
"order_id": "string", // 订单id 
"create_time": "string", // 参与时间 
"update_time": "string", //  
 */

type LotteryUser struct {
    Id int `db:"id" json:"id"`
    UserId string `db:"user_id" json:"user_id"`
    Mobile string `db:"mobile" json:"mobile"`
    LotteryId int `db:"lottery_id" json:"lottery_id"`
    LotteryAwardId int `db:"lottery_award_id" json:"lottery_award_id"`
    UseGiftsSn string `db:"use_gifts_sn" json:"use_gifts_sn"`
    CostIntegral int `db:"cost_integral" json:"cost_integral"`
    Status string `db:"status" json:"status"`
    OrderId string `db:"order_id" json:"order_id"`
    CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
    UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewLotteryUserModel() *LotteryUser {
	return &LotteryUser{}
}

func (m *LotteryUser) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *LotteryUser) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *LotteryUser) Update(arg *LotteryUser) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *LotteryUser) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}


func (m *LotteryUser) SetId(arg int) *LotteryUser {
	m.Id = arg
	return m
}

func (m *LotteryUser) SetUserId(arg string) *LotteryUser {
	m.UserId = arg
	return m
}

func (m *LotteryUser) SetMobile(arg string) *LotteryUser {
	m.Mobile = arg
	return m
}

func (m *LotteryUser) SetLotteryId(arg int) *LotteryUser {
	m.LotteryId = arg
	return m
}

func (m *LotteryUser) SetLotteryAwardId(arg int) *LotteryUser {
	m.LotteryAwardId = arg
	return m
}

func (m *LotteryUser) SetUseGiftsSn(arg string) *LotteryUser {
	m.UseGiftsSn = arg
	return m
}

func (m *LotteryUser) SetCostIntegral(arg int) *LotteryUser {
	m.CostIntegral = arg
	return m
}

func (m *LotteryUser) SetStatus(arg string) *LotteryUser {
	m.Status = arg
	return m
}

func (m *LotteryUser) SetOrderId(arg string) *LotteryUser {
	m.OrderId = arg
	return m
}

func (m *LotteryUser) SetCreateTime(arg string) *LotteryUser {
	m.CreateTime = arg
	return m
}

func (m *LotteryUser) SetUpdateTime(arg string) *LotteryUser {
	m.UpdateTime = arg
	return m
}

func (m LotteryUser) AsMapItf() MapItf {
	return MapItf{ 
        "id": m.Id, 
        "user_id": m.UserId, 
        "mobile": m.Mobile, 
        "lottery_id": m.LotteryId, 
        "lottery_award_id": m.LotteryAwardId, 
        "use_gifts_sn": m.UseGiftsSn, 
        "cost_integral": m.CostIntegral, 
        "status": m.Status, 
        "order_id": m.OrderId, 
        "create_time": m.CreateTime, 
        "update_time": m.UpdateTime, 
	}
}
func (m LotteryUser) Translates() map[string]string {
	return map[string]string{ 
        "id": "", 
        "user_id": "用户编号", 
        "mobile": "手机号码", 
        "lottery_id": "活动编号", 
        "lottery_award_id": "奖品编号", 
        "use_gifts_sn": "兑奖号码", 
        "cost_integral": "消耗的积分", 
        "status": "发奖状态 INIT 未发奖 DONE 已发奖", 
        "order_id": "订单id", 
        "create_time": "参与时间", 
        "update_time": "", 
	}
}