package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
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
