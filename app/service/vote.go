package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
)

type VoteService struct{}

func (VoteService) CreateVote(bustimeId int64, userId int64, previous bool, nearest bool, next bool) error {
	vote := model.Vote{
		BusTimeId: bustimeId,
		UserId:    userId,
		Previous:  previous,
		Nearest:   nearest,
		Next:      next,
	}
	if err := lib.DB.Create(&vote).Error; err != nil {
		return err
	}
	return nil
}
