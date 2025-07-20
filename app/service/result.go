package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"log"
)

type ResultService struct{}

func (ResultService) GetResult(busTimeId int64) ([]model.Result, error) {
	results := []model.Result{}

	query := lib.DB.Model(&model.Result{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}

	if err := query.Find(&results).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return results, nil
}
