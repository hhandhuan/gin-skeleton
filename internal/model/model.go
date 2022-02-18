package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type gTime struct {
	time.Time
}

// MarshalJSON on gTime format Time field with %Y-%m-%d %H:%M:%S
func (t gTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t gTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *gTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = gTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
