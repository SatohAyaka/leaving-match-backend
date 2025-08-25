package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"time"
)

type RecommendedService struct{}

func (RecommendedService) CreateRecommended(recommendedTime time.Time, memberIds []int, status bool) (model.Recommended, error) {
	recommended := model.Recommended{
		RecommendedTime: recommendedTime,
		MemberIds:       memberIds,
		Status:          status,
	}
	if err := lib.DB.Create(&recommended).Error; err != nil {
		return model.Recommended{}, err
	}

	return recommended, nil
}
