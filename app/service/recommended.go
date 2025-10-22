package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"time"
)

type RecommendedService struct{}

func (RecommendedService) CreateRecommended(recommendedTime time.Time, memberIds []int64) (model.RecommendedResponse, error) {
	now := time.Now()
	status := false
	if diff := recommendedTime.Sub(now); diff > 10 && diff <= 50*time.Minute {
		status = true
	}
	recommended := model.Recommended{
		RecommendedTime: recommendedTime,
		MemberIds:       memberIds,
		Status:          status,
	}
	if err := lib.DB.Create(&recommended).Error; err != nil {
		return model.RecommendedResponse{}, err
	}

	return model.RecommendedResponse{
		RecommendedId: recommended.RecommendedId,
		Status:        recommended.Status,
	}, nil
}

func (RecommendedService) GetLatestRecommended() (model.Recommended, error) {
	var latest model.Recommended
	if err := lib.DB.Order("created_date DESC").First(&latest).Error; err != nil {
		return model.Recommended{}, err
	}
	return latest, nil
}
