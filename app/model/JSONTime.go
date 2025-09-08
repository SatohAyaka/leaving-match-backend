package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Format("2006-01-02 15:04:05"))
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Time = time.Time{}
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsed, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

func (t JSONTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		t.Time = v
	case []byte:
		parsed, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		t.Time = parsed
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}
