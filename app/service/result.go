package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"log"
	"time"
)

type ResultService struct{}

func (ResultService) CreateResult(busTimeId int64, busTime time.Time, member int64) (int64, error) {
	results := model.Result{
		BusTimeId: busTimeId,
		BusTime:   model.JSONTime{Time: busTime},
		Member:    member,
	}
	if err := lib.DB.Create(&results).Error; err != nil {
		return 0, err
	}
	return results.ResultId, nil
}

func (ResultService) GetResult(busTimeId int64) (model.Result, error) {
	var result model.Result

	query := lib.DB.Model(&model.Result{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}

	if err := query.Find(&result).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return model.Result{}, err
	}

	return result, nil
}

func (ResultService) GetLatestResult() (model.Result, error) {
	var result model.Result
	if err := lib.DB.Order("created_date DESC").First(&result).Error; err != nil {
		return model.Result{}, err
	}
	return result, nil
}
