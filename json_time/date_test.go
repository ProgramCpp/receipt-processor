package json_time_test

import (
	"testing"

	"github.com/programcpp/receipt-processor/json_time"
	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T){
	date := json_time.Date{}
	err := date.UnmarshalJSON([]byte("\"2023-11-01\""))
	assert.NoError(t, err)
	
	dateStr, err := date.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "\"2023-11-01\"", string(dateStr))
}