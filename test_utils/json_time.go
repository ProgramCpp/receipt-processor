package test_utils

import (
	"fmt"

	"github.com/programcpp/receipt-processor/json_time"
)

func GetDate(date string) json_time.Date {
	d := json_time.Date{}
	_ = d.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", date)))
	return d
}

func GetTime(time string) json_time.Time {
	t := json_time.Time{}
	_ = t.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", time)))
	return t
}