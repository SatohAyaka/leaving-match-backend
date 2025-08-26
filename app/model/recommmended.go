package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type IntSlice []int64

func (s IntSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *IntSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan IntSlice: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

type Recommended struct {
	RecommendedId   int64     `gorm:"column:recommended_id;primaryKey;autoIncrement"`
	RecommendedTime time.Time `gorm:"column:recommended_time"`
	MemberIds       IntSlice  `gorm:"column:member_ids;type:json;not null"`
	Status          bool      `gorm:"column:status"`
	CreatedDate     time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

type RecommendedResponse struct {
	RecommendedId int64
	Status        bool
}

func (Recommended) TableName() string { return "Recommended_Data" }
