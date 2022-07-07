package service

import (
	"active/app/dao"
	"active/app/model"
	. "active/common"
	"math/rand"
	"time"
)

type LotterySrv struct{}

func NewLotterySrv() *LotterySrv {
	return &LotterySrv{}
}

// 获取抽奖详情
type InfoLotteryArgs struct {
	Code string `json:"code"`
}

type InfoLotteryRet struct {
	Info       *model.Lottery       `json:"info"`
	AwardCount int                  `json:"award_count"`
	AwardList  []model.LotteryAward `json:"award_list"`
}

func (s *LotterySrv) InfoLottery(args *InfoLotteryArgs) (data InfoLotteryRet) {
	Lottery := &model.Lottery{
		Code: args.Code,
	}
	if !Lottery.Info() {
		return
	}
	data.Info = Lottery
	data.AwardCount, data.AwardList = dao.ListLotteryAward(Lottery.Id)
	return
}

// 新增抽奖记录
type InsertLotteryUserArgs struct {
	UserId    string `json:"user_id"`
	LotteryId int    `json:"lottery_id"`
}

type InsertLotteryUserRet struct {
	LotteryAwardId int `json:"lottery_award_id"`
	UseGiftsSn string `json:"use_gifts_sn"`
	AwardType string `json:"award_type"`
	Img string `json:"img"`
}

func (s *LotterySrv) InsertLotteryUser(args *InsertLotteryUserArgs) (data InsertLotteryUserRet, respErr RespErr) {
	lottery := &model.Lottery{
		Id: args.LotteryId,
	}
	if !lottery.Info() {
		return data, NewRespErr(ErrValidReq, "无效抽奖活动")
	}

	//判断当天的抽奖次数是否已经超过限制
	dayLotteryCount := dao.GetLotteryUserCnt(&dao.GetLotteryUserCntArgs{
		LotteryId: args.LotteryId,
		UserId: args.UserId,
		StartTime: time.Now().Format("2006-01-02" + " 00:00:00"),
		EndTime: time.Now().Format("2006-01-02") + " 23:59:59",
	})
	if lottery.DayLotteryLimit > 0 && dayLotteryCount >= lottery.DayLotteryLimit {
		return data, NewRespErr(ErrActiveJoinCount, "今日抽奖次数用完")
	}

	winCnt, _ := dao.GetLotteryUser(&dao.GetLotteryUserArgs{
		UserId: args.UserId,
		LotteryId: args.LotteryId,
	})

	winAward := model.LotteryAward{}
	if winCnt > 0 {
		winAward.Id = 8
		winAward.Info()
	} else {
		// 抽奖
		winAward, respErr = calAwardNew(lottery.Id)
		if respErr.NotNil() {
			return
		}
	}
	useGiftsSn := RandomStr(4)

	lotteryUserSet := &model.LotteryUser{
		UserId:         args.UserId,
		LotteryId:      lottery.Id,
		LotteryAwardId: winAward.Id,
		Status:         "INIT",
		UseGiftsSn:     useGiftsSn,
	}
	row := lotteryUserSet.Insert()
	if row == 0 {
		return data, NewRespErr(ErrInsert, "抽奖失败")
	}

	// 已中奖品数量 +1
	dao.LotteryAwardIncrWinnerNum(winAward.Id)

	data.LotteryAwardId = winAward.Id
	data.UseGiftsSn = useGiftsSn
	data.AwardType = winAward.Type
	data.Img = winAward.Img
	return
}

// 获取中奖记录列表
func (s *LotterySrv) ListLotteryUser(args *dao.ListLotteryUserArgs) (data []dao.ListLotteryUserRet, respErr RespErr) {
	_, data = dao.ListLotteryUser(args)
	return
}

// 获取用户中奖记录
func (s *LotterySrv) GetLotteryUser(args *dao.GetLotteryUserArgs) (data []dao.GetLotteryUserRet, respErr RespErr) {
	_, data = dao.GetLotteryUser(args)
	return
}

//计算奖品（固定概率）
func calAwardNew(lotteryId int) (winAward model.LotteryAward, respErr RespErr) {
	_, awardList := dao.ListLotteryAward(lotteryId)
	awardMap := make(map[int]model.LotteryAward)
	awardIdList := make([]int, 0)
	for _, award := range awardList {
		awardMap[award.Id] = award
		num := award.TotalNum - award.WinnerNum
		if num < 1 {
			//准备的奖项已经抽完,使用备用
			num = award.ReserveNum - award.ReserveWinnerNum
		}
		//将奖项写入数组
		for i := 0; i < num; i++ {
			awardIdList = append(awardIdList, award.Id)
		}
	}
	if len(awardIdList) == 0 {
		respErr = NewRespErr(ErrActiveLotteryAward, "")
	}
	awardIdList = Shuffle(awardIdList)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key := r.Intn(len(awardIdList))

	idKey := awardIdList[key]
	winAward = awardMap[idKey]
	return
}

func Shuffle(array []int) []int {
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	in := rr.Perm(len(array))

	out := make([]int, 0, len(array))
	for _, v := range in {
		out = append(out, array[v])
	}
	return out
}
