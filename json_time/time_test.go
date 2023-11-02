package json_time_test

import (
	"testing"

	"github.com/programcpp/receipt-processor/json_time"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T){
	time := json_time.Time{}
	err := time.UnmarshalJSON([]byte("\"23:30\""))
	assert.NoError(t, err)
	

	dateStr, err := time.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "\"23:30\"", string(dateStr))
}

func TestTimeBadFormat(t *testing.T){
	time := json_time.Time{}
	err := time.UnmarshalJSON([]byte("\"2023-11-01 23:30:30\""))
	assert.Error(t, err)

	err = time.UnmarshalJSON([]byte("\"2006-01-02T15:04:05Z07:00\""))
	assert.Error(t, err)

	err = time.UnmarshalJSON([]byte("\"11:30PM\""))
	assert.Error(t, err)

	err = time.UnmarshalJSON([]byte("\"23:30:00\""))
	assert.Error(t, err)

	err = time.UnmarshalJSON([]byte("\"\""))
	assert.Error(t, err)
}