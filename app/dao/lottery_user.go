package dao

import (
	. "active/common"
)

// 获取用户中奖记录
type GetLotteryUserArgs struct {
	UserId    string `json:"user_id"`
	LotteryId int    `json:"lottery_id"`
}

type GetLotteryUserRet struct {
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	Img        string `json:"img"`
	UseGiftsSn string `json:"use_gifts_sn"`
}

func GetLotteryUser(args *GetLotteryUserArgs) (int, []GetLotteryUserRet) {
	var lotteryUserRet []GetLotteryUserRet
	count, err := GetDbEngineIns().Table("my_lottery_user").Alias("lu").
		Select("la.name, la.img, lu.create_time, lu.use_gifts_sn").
		Join("left", []string{"my_lottery_award", "la"}, "lu.lottery_award_id = la.id").
		Where("lu.lottery_id = ?", args.LotteryId).
		And("lu.user_id = ?", args.UserId).
		And("la.type != 'none'").
		FindAndCount(&lotteryUserRet)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), lotteryUserRet
}

// 获取中奖记录列表
type ListLotteryUserArgs struct {
	LotteryId int `json:"lottery_id"`
}
type ListLotteryUserRet struct {
	OpenId     string `json:"open_id"`
	NickName   string `json:"nick_name"`
	AwardName  string `json:"award_name"`
	AwardType  string `json:"award_type"`
	CreateTime string `json:"create_time"`
	UseGiftsSn string `json:"use_gifts_sn"`
}

func ListLotteryUser(args *ListLotteryUserArgs) (int, []ListLotteryUserRet) {
	var lotteryUserRet []ListLotteryUserRet
	count, err := GetDbEngineIns().Table("my_lottery_user").Alias("lu").
		Select("u.open_id, u.nick_name, la.name as award_name, lu.create_time, lu.use_gifts_sn, la.type award_type").
		Join("left", []string{"my_lottery_award", "la"}, "lu.lottery_award_id = la.id").
		Join("left", []string{"my_user", "u"}, "u.id = lu.user_id").
		Where("lu.lottery_id = ?", args.LotteryId).
		OrderBy("lu.create_time desc").
		FindAndCount(&lotteryUserRet)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), lotteryUserRet
}

//  获取用户一段时间内抽奖次数
type GetLotteryUserCntArgs struct {
	UserId    string `json:"user_id"`
	LotteryId int    `json:"lottery_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func GetLotteryUserCnt(args *GetLotteryUserCntArgs) int {
	count, err := GetDbEngineIns().Table("my_lottery_user").
		Where("lottery_id = ?", args.LotteryId).
		And("user_id = ?", args.UserId).
		And("create_time >= ?", args.StartTime).
		And("create_time <= ?", args.EndTime).
		Count()
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count)
}
