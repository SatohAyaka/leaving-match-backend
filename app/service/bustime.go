package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"log"
	"time"
)

type BusTimeService struct{}

func (BusTimeService) CreateBusTime(member string, previous time.Time, nearest time.Time, next time.Time) (int64, error) {
	bustime := model.BusTime{
		MemberId:     member,
		PreviousTime: previous,
		NearestTime:  nearest,
		NextTime:     next,
	}

	if err := lib.DB.Create(&bustime).Error; err != nil {
		return 0, err
	}

	return bustime.BusTimeId, nil
}
func (BusTimeService) GetBusTime(busTimeId int64) ([]model.BusTime, error) {
	bustimeData := []model.BusTime{}

	query := lib.DB.Model(&model.BusTime{})
	if busTimeId > 0 {
		query = query.Where("BusTimeId = ?", busTimeId)
	}

	if err := query.Find(&bustimeData).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return bustimeData, nil
}
