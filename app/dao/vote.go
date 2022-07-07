package dao

import (
	"active/app/model"
	. "active/common"
)

func VoteIncJoinTime(id int) {
	_, err := GetDbEngineIns().Table("my_vote").
		Incr("join_times").
		Where("id = ?", id).
		Update(&model.Vote{})
	if err != nil {
		panic(NewDbErr(err))
	}
}

func VoteIncVisitTime(id int) {
	_, err := GetDbEngineIns().Table("my_vote").
		Where("id = ?", id).
		Incr("visits_times").Update(&model.Vote{})
	if err != nil {
		panic(NewDbErr(err))
	}
}