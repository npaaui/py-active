package service

import (
	"active/app/dao"
	"active/app/model"
	. "active/common"
	"time"
)

type VoteSrv struct{}

func NewVoteSrv() *VoteSrv {
	return &VoteSrv{}
}

// 获取投票详情
type InfoVoteArgs struct {
	Code string `json:"code"`
}

type InfoVoteRet struct {
	Info *model.Vote `json:"info"`
	OptionCount int `json:"option_count"`
	OptionList []model.VoteOption `json:"option_list"`
}

func (s *VoteSrv) InfoVote(args *InfoVoteArgs) (data InfoVoteRet) {
	vote := &model.Vote{
		Code: args.Code,
	}
	if !vote.Info() {
		return
	}
	data.Info = vote
	data.OptionCount, data.OptionList = dao.ListVoteOption(vote.Id)

	// 访问次数+1
	dao.VoteIncVisitTime(vote.Id)
	return
}

// 新增投票记录
type InsertVoteUserArgs struct {
	UserId string `json:"user_id"`
	VoteId 	int `json:"vote_id"`
	Content map[int]string `json:"code"`
}

func (s *VoteSrv) InsertVoteUser(args *InsertVoteUserArgs) (respErr RespErr) {
	vote := &model.Vote{
		Id: args.VoteId,
	}
	if !vote.Info() {
		return NewRespErr(ErrValidReq, "无效投票")
	}
	if vote.EndTime < time.Now().Format("2006-01-02 15:04:05") {
		return NewRespErr(ErrValidReq, "活动已结束")
	}

	//判断当天的投票次数是否已经超过限制
	dayVoteCount := dao.GetVoteUserCnt(&dao.GetVoteUserCntArgs{
		VoteId: args.VoteId,
		UserId: args.UserId,
		StartTime: time.Now().Format("2006-01-02" + " 00:00:00"),
		EndTime: time.Now().Format("2006-01-02") + " 23:59:59",
	})
	if vote.DayVoteLimit > 0 && dayVoteCount >= vote.DayVoteLimit {
		return NewRespErr(ErrActiveJoinCount, "今日投票次数用完")
	}

	optMap := map[int]model.VoteOption{}
	_, options := dao.ListVoteOption(vote.Id)
	for _, option := range options {
		optMap[option.Id] = option
	}

	voteTime := time.Now().Format("2006-01-02 15:04:05")
	for k, v := range args.Content {
		if optMap[k].Type == "check" && v != "1" && v != "0" {
			return NewRespErr(ErrValidReq, "选项" + IntToStr(k) + "结果有误")
		}

		optSet := &model.VoteUser{
			UserId: args.UserId,
			VoteId: vote.Id,
			VoteOptionId: k,
			Value: v,
			VoteTime: voteTime,
		}
		row := optSet.Insert()
		if row == 0 {
			return NewRespErr(ErrInsert, "选项" + IntToStr(k) + "结果记录失败")
		}
		dao.VoteOptionIncVotes(k)
	}

	// 参与次数+1
	dao.VoteIncJoinTime(vote.Id)
	return
}
