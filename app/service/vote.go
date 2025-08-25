package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"log"
)

type VoteService struct{}

func (VoteService) CreateVote(bustimeId int64, userId int64, previous bool, nearest bool, next bool) error {
	vote := model.Vote{
		BusTimeId:     bustimeId,
		BackendUserId: userId,
		Previous:      previous,
		Nearest:       nearest,
		Next:          next,
	}
	if err := lib.DB.Create(&vote).Error; err != nil {
		return err
	}
	return nil
}

func (VoteService) GetVote(busTimeId int64) ([]model.Vote, error) {
	allvote := []model.Vote{}

	query := lib.DB.Model(&model.Vote{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}
	if err := query.Find(&allvote).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	// ユーザごとの最新投票取得
	userVotes := make(map[int64]model.Vote)
	for _, vote := range allvote {
		existing, ok := userVotes[vote.BackendUserId]
		if !ok || vote.VoteId > existing.VoteId {
			userVotes[vote.BackendUserId] = vote
		}
	}

	// ↑から一つのVoteデータに
	votes := make([]model.Vote, 0, len(userVotes))
	for _, vote := range userVotes {
		votes = append(votes, vote)
	}

	return votes, nil
}
