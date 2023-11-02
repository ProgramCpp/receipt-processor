package points

import (
	"math"
	"strings"
	"time"

	"github.com/programcpp/receipt-processor/receipts"
)

// all the rules can be independently implemented and independently tested
// func type is sufficient in this case. rules can be implemented as functions. interface introduces struct bloat
type Rule func(receipts.Receipt) int

// One point for every alphanumeric character in the retailer name
func RetailerRule(r receipts.Receipt) int {
	return len(r.Retailer)
}

// 50 points if the total is a round dollar amount with no cents
func RoundTotalRule(r receipts.Receipt) int {
	if math.Trunc(r.Total) == r.Total {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25
func MultipleTotalRule(r receipts.Receipt) int {
	if math.Mod(r.Total, 0.25) == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt.
func ItemCountRule(r receipts.Receipt) int {
	return 5 * (len(r.Items) / 2)
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func ItemDescriptionRule(r receipts.Receipt) int {
	points := 0
	for _, i := range r.Items {
		points += ItemPoints(i)
	}
	return points
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func ItemPoints(i receipts.Item) int {
	if len(strings.TrimSpace(i.ShortDescription))%3 == 0 {
		return int(math.Ceil(i.Price * 0.2))
	}
	return 0
}

// 6 points if the day in the purchase date is odd.
func PurchaseDateRule(r receipts.Receipt) int {
	if r.PurchaseDate.Day()%2 == 1 {
		return 6
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm. 
// () exclusive! for simplicity
func PurchaseTimeRule(r receipts.Receipt) int {
	// the input is parsed from string. for equality, base time must be parsed too
	// time struct have different default values. https://github.com/golang/go/issues/40925#issuecomment-677775168
	t2PM, _ := time.Parse(time.TimeOnly, "14:00:00") // parsing hardcoded value will succeed!
	t4PM, _ := time.Parse(time.TimeOnly, "16:00:00") // parsing hardcoded value will succeed!
	if r.PurchaseTime.After(t2PM) && r.PurchaseTime.Before(t4PM) {
		return 10
	}
	return 0
}
