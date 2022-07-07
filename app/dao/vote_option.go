package dao

import (
	"active/app/model"
	. "active/common"
)

func ListVoteOption(id int) (int, []model.VoteOption) {
	var optList []model.VoteOption
	count, err := GetDbEngineIns().Table("my_vote_option").
		Where("vote_id = ?", id).FindAndCount(&optList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), optList
}

func VoteOptionIncVotes(id int) {
	_, err := GetDbEngineIns().Table("my_vote_option").
		Where("id = ?", id).
		Incr("votes").
		Update(model.VoteOption{})
	if err != nil {
		panic(NewDbErr(err))
	}
}
