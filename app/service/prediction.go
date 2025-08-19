package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"log"
	"time"
)

type PredictionService struct{}

func (PredictionService) CreatePrediction(predictionId int64, userId int64, predictionTime time.Time) error {
	prediction := model.Prediction{
		PredictionId:   predictionId,
		UserId:         userId,
		PredictionTime: predictionTime,
	}

	if err := lib.DB.Create(&prediction).Error; err != nil {
		return err
	}

	return nil
}

func (PredictionService) GetPrediction(busTimeId int64) ([]model.Prediction, error) {
	predictions := []model.Prediction{}

	query := lib.DB.Model(&model.Prediction{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}
	if err := query.Find(&predictions).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}
	return predictions, nil
}
