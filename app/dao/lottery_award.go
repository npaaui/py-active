package dao

import (
	"active/app/model"
	. "active/common"
)

func ListLotteryAward(id int) (int, []model.LotteryAward) {
	var awardList []model.LotteryAward
	count, err := GetDbEngineIns().Table("my_lottery_award").
		Where("lottery_id = ?", id).
		OrderBy("display_order asc").
		FindAndCount(&awardList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), awardList
}

func LotteryAwardIncrWinnerNum(id int) {
	_, err := GetDbEngineIns().Table("my_lottery_award").
		Where("id = ?", id).
		Incr("winner_num").
		Update(model.LotteryAward{})
	if err != nil {
		panic(NewDbErr(err))
	}
}

