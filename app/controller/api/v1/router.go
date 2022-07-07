package v1

import (
	"active/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	voteCtr := NewVoteCtr()
	lotteryCtr := NewLotteryCtr()
	userCtr := NewUserCtr()
	configCtr := NewConfigCtr()

	r.Use(middleware.ReqLog(), middleware.RecoverDbError(), middleware.Access())

	v1 := r.Group("/api/v1")
	{
		v1.POST("user", userCtr.InsertUser)
		v1.GET("vote", voteCtr.InfoVote)
		v1.GET("lottery", lotteryCtr.InfoLottery)
		v1.GET("config", configCtr.InfoConfig)
		v1.GET("vote_option", voteCtr.ListVoteOption)
		v1.GET("lottery_user/list", lotteryCtr.ListLotteryUser)
	}

	auth := r.Group("/api/v1").Use(middleware.AuthUser())
	{
		auth.POST("vote_user", voteCtr.InsertVoteUser)
		auth.POST("lottery_user", lotteryCtr.InsertLotteryUser)
		auth.GET("lottery_user", lotteryCtr.GetLotteryUser)
	}
}
