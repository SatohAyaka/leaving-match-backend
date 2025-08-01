package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"fmt"
	"log"
	"time"
)

type BusTimeService struct{}

func (BusTimeService) CreateBusTime(previous time.Time, nearest time.Time, next time.Time) (int64, error) {
	bustime := model.BusTime{
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
	bustimes := []model.BusTime{}

	query := lib.DB.Model(&model.BusTime{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}

	if err := query.Find(&bustimes).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return bustimes, nil
}

func (BusTimeService) BusTimeToId(busTimeId int64, selectTime int64) (time.Time, error) {
	bustimes := model.BusTime{}

	query := lib.DB.Model(&model.BusTime{})
	if busTimeId > 0 {
		query = query.Where("bustime_id = ?", busTimeId)
	}

	if err := query.First(&bustimes).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return time.Time{}, err
	}

	switch selectTime {
	case 1:
		return bustimes.PreviousTime, nil
	case 2:
		return bustimes.NearestTime, nil
	case 3:
		return bustimes.NextTime, nil
	default:
		return time.Time{}, fmt.Errorf("invalid selectTime: %d", selectTime)
	}
}
