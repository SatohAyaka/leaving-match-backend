package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"log"
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

func (PredictionService) GetPrediction(busTimeId int64) ([]model.Prediction, error) {
	predictionData := []model.Prediction{}

	query := lib.DB.Model(&model.Prediction{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}
	if err := query.Find(&predictionData).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}
	return predictionData, nil
}
