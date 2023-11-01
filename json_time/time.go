package json_time

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
type to support custom time format json codec
*/

const TIME_FORMAT = "15:04"

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var timeStr string
	err := json.Unmarshal(b, &timeStr)
	if err != nil {
		return nil
	}
	
	parsed, err := time.Parse(TIME_FORMAT, timeStr)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}


func (t *Time) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"",t.Format(TIME_FORMAT))
	return []byte(s), nil
}