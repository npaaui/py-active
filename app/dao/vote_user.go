package dao

import (
	. "active/common"
)

// 获取用户中奖记录
type GetVoteUserArgs struct {
	UserId string `json:"user_id"`
	VoteId int    `json:"vote_id"`
}

type GetVoteUserRet struct {
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	Img        string `json:"img"`
	UseGiftsSn string `json:"use_gifts_sn"`
}

func GetVoteUser(args *GetVoteUserArgs) (int, []GetVoteUserRet) {
	var voteUserRet []GetVoteUserRet
	count, err := GetDbEngineIns().Table("my_vote_user").Alias("lu").
		Select("la.name, la.img, lu.create_time, lu.use_gifts_sn").
		Join("left", []string{"my_vote_award", "la"}, "lu.vote_award_id = la.id").
		Where("lu.vote_id = ?", args.VoteId).
		And("lu.user_id = ?", args.UserId).
		And("la.type != 'none'").
		FindAndCount(&voteUserRet)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), voteUserRet
}

//  获取用户一段时间内抽奖次数
type GetVoteUserCntArgs struct {
	UserId    string `json:"user_id"`
	VoteId    int    `json:"vote_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func GetVoteUserCnt(args *GetVoteUserCntArgs) int {
	list, err := GetDbEngineIns().Table("my_vote_user").
		Select("user_id, vote_id, vote_time").
		Where("user_id = ?", args.UserId).
		And("create_time >= ?", args.StartTime).
		And("create_time <= ?", args.EndTime).
		GroupBy("user_id,vote_id,vote_time").
		QueryString()
	if err != nil {
		panic(NewDbErr(err))
	}
	return len(list)
}
