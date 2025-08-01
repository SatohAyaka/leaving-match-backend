package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"time"
)

type PredictionService struct{}

func (PredictionService) CreatePrediction(busTimeId int64, userId int64, predictionTime time.Time) error {
	prediction := model.Prediction{
		BusTimeId:      busTimeId,
		UserId:         userId,
		PredictionTime: predictionTime,
	}

	if err := lib.DB.Create(&prediction).Error; err != nil {
		return err
	}

	return nil
}
