package json_time

import "time"

/*
type to support custom time format json codec
*/

const TIME_FORMAT = "15:04"

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(TIME_FORMAT, string(b))
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}


func (t *Time) MarshalJSON() ([]byte, error) {
	s := t.Format(TIME_FORMAT)
	return []byte(s), nil
}