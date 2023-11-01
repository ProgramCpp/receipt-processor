package json_time

import (
	"time"
)

/*
type to support custom date format json codec
*/

const DATE_FORMAT = time.DateOnly

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(DATE_FORMAT, string(b))
	if err != nil {
		return err
	}

	d.Time = parsed
	return nil
}


func (d *Date) MarshalJSON() ([]byte, error) {
	s := d.Format(DATE_FORMAT)
	return []byte(s), nil
}