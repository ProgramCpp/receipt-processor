package json_time_test

import (
	"testing"

	"github.com/programcpp/receipt-processor/utils/json_time"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T){
	date := json_time.Time{}
	err := date.UnmarshalJSON([]byte("23:30"))
	assert.NoError(t, err)
	

	dateStr, err := date.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "23:30", string(dateStr))
}