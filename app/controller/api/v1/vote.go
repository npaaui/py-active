package v1

import (
	"active/app/dao"
	"active/app/model"
	"active/app/service"
	. "active/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type VoteCtr struct {
	Srv *service.VoteSrv
}

func NewVoteCtr() *VoteCtr {
	return &VoteCtr{
		Srv: service.NewVoteSrv(),
	}
}

func (t *VoteCtr) InfoVote(c *gin.Context) {
	args := &service.InfoVoteArgs{}
	ValidateQuery(c, map[string]string{
		"code": "string|required|enum:py_food",
	}, args)

	vote := t.Srv.InfoVote(args)
	if vote.Info == nil {
		ReturnErrMsg(c, ErrNotExist, "无效投票")
		return
	}
	startTime, _ := time.Parse("2006-01-02 15:04:05", vote.Info.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", vote.Info.EndTime)
	vote.Info.StartTime = startTime.Format("2006-01-02 15:04")
	vote.Info.EndTime = endTime.Format("2006-01-02 15:04")

	checkCnt := 0
	optionList := map[string][]model.VoteOption{}
	for _, v := range vote.OptionList {
		if v.Type == "check" {
			checkCnt++
		}
		optionList[v.Group] = append(optionList[v.Group], v)
	}

	// 统计字段
	type DataStatistic struct {
		Name string `json:"name"`
		Num  int    `json:"num"`
		Unit string `json:"unit"`
	}
	// 选项数据
	type Option struct {
		Title string `json:"title"`
		MaxSelect int `json:"max_select"`
		MinSelect int `json:"min_select"`
		List []model.VoteOption `json:"list"`
	}
	// 奖品数据
	type Prize struct {
		Level string `json:"level"`
		Title string `json:"title"`
		Img string `json:"img"`
	}
	_, award := dao.ListLotteryAward(1)

	data := struct {
		Info *model.Vote `json:"info"`
		OptionList []Option `json:"option_list"`
		DataShow []DataStatistic `json:"data_show"`
		PrizeList []Prize `json:"prize_list"`
	}{
		Info: vote.Info,
		OptionList: []Option{
			{
				Title: optionList["food"][0].Remark,
				MaxSelect: vote.Info.MaxSelect,
				MinSelect: vote.Info.MinSelect,
				List: optionList["food"],
			},
			{
				Title: optionList["snack"][0].Remark,
				MaxSelect: vote.Info.MaxSelect,
				MinSelect: vote.Info.MinSelect,
				List: optionList["snack"],
			},
		},
		DataShow: []DataStatistic{
			{
				Name: "菜品数量",
				Num: checkCnt,
				Unit: "道",
			},{
				Name: "投票次数",
				Num: vote.Info.JoinTimes,
				Unit: "次",
			},{
				Name: "访问次数",
				Num: vote.Info.VisitsTimes,
				Unit: "次",
			},
		},
		PrizeList: []Prize{
			{
				Level: "一等奖",
				Title: award[0].Name,
				Img: award[0].Img,
			},
			{
				Level: "二等奖",
				Title: award[1].Name,
				Img: award[1].Img,
			},
			{
				Level: "三等奖",
				Title: award[2].Name,
				Img: award[2].Img,
			},
		},
	}

	ReturnData(c, data)
}

func (t *VoteCtr) InsertVoteUser(c *gin.Context) {
	args := &struct {
		VoteId  int `json:"vote_id"`
		Content string `json:"content"`
	}{}
	_ = ValidatePostJson(c, map[string]string{
		"vote_id": "int|required",
		"content": "string|required",
	}, args)

	content := map[int]string{}
	err := json.Unmarshal([]byte(args.Content), &content)
	if err != nil {
		ReturnErrMsg(c, ErrValidReq, "无效投票结果")
		return
	}
	respErr := t.Srv.InsertVoteUser(&service.InsertVoteUserArgs{
		UserId:  c.GetString("user_id"),
		VoteId:  args.VoteId,
		Content: content,
	})
	if respErr.NotNil() {
		ReturnByRespErr(c, respErr)
		return
	}

	ReturnData(c, nil)
}

func (t *VoteCtr) ListVoteOption(c *gin.Context) {
	ReturnData(c, nil)
}
