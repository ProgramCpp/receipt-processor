package points

import (
	"github.com/programcpp/receipt-processor/receipts"
)

// rule engine to run all the rules
type Engine struct{
	rules []Rule
}

// init all the rules
// TODO: inject rules for testing the engine. add tests for the engine
func NewEngine() Engine{
	return Engine{
		rules: []Rule{},
	}
}

func (e Engine)Points(r receipts.Receipt) int {
	points := 0

	for _, rule := range e.rules {
		points += rule(r)
	}

	return points
}




