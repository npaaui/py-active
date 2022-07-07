package v1

import (
	"active/app/dao"
	"active/app/service"
	. "active/common"
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"time"
)

type LotteryCtr struct {
	Srv *service.LotterySrv
}

func NewLotteryCtr() *LotteryCtr {
	return &LotteryCtr{
		Srv: service.NewLotterySrv(),
	}
}

func (t *LotteryCtr) InfoLottery(c *gin.Context) {
	args := &service.InfoLotteryArgs{}
	ValidateQuery(c, map[string]string{
		"code": "string|required|enum:py_food",
	}, args)

	Lottery := t.Srv.InfoLottery(args)
	if Lottery.Info == nil {
		ReturnErrMsg(c, ErrNotExist, "无效抽奖活动")
		return
	}

	ReturnData(c, Lottery)
	return
}

func (t *LotteryCtr) InsertLotteryUser(c *gin.Context) {
	args := &service.InsertLotteryUserArgs{
		UserId:    c.GetString("user_id"),
		LotteryId: 0,
	}
	_ = ValidatePostJson(c, map[string]string{
		"lottery_id": "int|required",
	}, args)

	ret, respErr := t.Srv.InsertLotteryUser(args)
	if respErr.NotNil() {
		ReturnByRespErr(c, respErr)
		return
	}

	ReturnData(c, ret)
}

func (t *LotteryCtr) GetLotteryUser(c *gin.Context) {
	args := &dao.GetLotteryUserArgs{
		UserId:    c.GetString("user_id"),
		LotteryId: 0,
	}
	ValidateQuery(c, map[string]string{
		"lottery_id": "int",
	}, args)

	if args.LotteryId == 0 {
		args.LotteryId = 1
	}
	ret, respErr := t.Srv.GetLotteryUser(args)
	if respErr.NotNil() {
		ReturnByRespErr(c, respErr)
		return
	}

	ReturnData(c, ret)
}

func (t *LotteryCtr) ListLotteryUser(c *gin.Context) {
	args := &dao.ListLotteryUserArgs{
		LotteryId: 0,
	}
	params := ValidateQuery(c, map[string]string{
		"lottery_id": "int",
		"import": "string",
	}, args)

	if args.LotteryId == 0 {
		args.LotteryId = 1
	}

	list, respErr := t.Srv.ListLotteryUser(args)
	if respErr.NotNil() {
		ReturnByRespErr(c, respErr)
		return
	}

	if tmp, ok := params["import"].(string); ok && tmp == "Y" {
		header := []string{"用户编号", "微信昵称", "中奖礼品", "兑换码", "抽奖时间"}

		b := &bytes.Buffer{}
		b.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，避免中文乱码
		wr := csv.NewWriter(b)
		_ = wr.Write(header) //按行shu

		var items []string
		for _, v := range list {
			userGiftSn := v.UseGiftsSn
			if v.AwardType == "none" {
				userGiftSn = ""
			}
			items = []string{
				v.OpenId,
				v.NickName,
				v.AwardName,
				userGiftSn,
				v.CreateTime,
			}
			_ = wr.Write(items)
		}
		wr.Flush()

		dateline := time.Now().Format("2006-01-02 15:04")
		ReturnFile(c, "寻味饶州截止至"+ dateline +"抽奖结果.csv", b.Bytes())
		return
	}
	return
}
