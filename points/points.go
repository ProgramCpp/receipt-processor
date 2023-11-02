package points

import (
	"math"
	"strings"
	"time"

	"github.com/programcpp/receipt-processor/receipts"
)



func Points(r receipts.Receipt) int {
	points := 0

	points += RetailerPoints(r.Retailer)

	// 50 points if the total is a round dollar amount with no cents
	if math.Trunc(r.Total) == r.Total {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(r.Total, 0.25) == 0 {
		points += 25
	}

	// 	5 points for every two items on the receipt.
	points += len(r.Items) / 2

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, i := range r.Items {
		points += ItemPoints(i)
	}

	// 6 points if the day in the purchase date is odd.
	if r.PurchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if r.PurchaseTime.After(time.Time{}.Add(14*time.Hour)) &&
		r.PurchaseTime.Before(time.Time{}.Add(16*time.Hour)) {
		points += 10
	}

	return points
}

// One point for every alphanumeric character in the retailer name
func RetailerPoints(r string) int{
	return len(r)
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func ItemPoints(i receipts.Item) int {
	points := 0
	if len(strings.TrimSpace(i.ShortDescription)) % 3 == 0 {
		points += int(math.Ceil(i.Price * 0.2))
	}
	return points
}
