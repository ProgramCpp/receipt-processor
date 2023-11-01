package json_time

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
type to support custom date format json codec. the json value is in string format.
format: \"2006-01-02\"
*/

const DATE_FORMAT = time.DateOnly

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var dateStr string 
	err := json.Unmarshal(b, &dateStr)
	if err != nil {
		return nil
	}
	parsed, err := time.Parse(DATE_FORMAT, dateStr)
	if err != nil {
		return err
	}

	d.Time = parsed
	return nil
}


func (d *Date) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"",d.Format(DATE_FORMAT))
	return []byte(s), nil
}