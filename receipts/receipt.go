package receipts

import (
	"fmt"
	"regexp"

	"github.com/programcpp/receipt-processor/json_time"
)

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}
type Receipt struct {
	Retailer     string         `json:"retailer"`
	PurchaseDate json_time.Date `json:"purchaseDate"`
	PurchaseTime json_time.Time `json:"purchaseTime"`
	Items        []Item         `json:"items"`
	Total        float64        `json:"total,string"`
}

func (i Item) IsValid() bool {
	isMatch, err := regexp.MatchString("^[\\w\\s\\-]+$", i.ShortDescription)
	if err != nil || !isMatch {
		return false
	}

	return validateFloat64(i.Price)
}

func (r Receipt) IsValid() bool {
	if r.Retailer == "" {
		return false
	}

	// no whitespaces
	isMatch, err := regexp.MatchString("^\\S+$", r.Retailer)
	if err != nil || !isMatch {
		return false
	}

	// invalid date and time formats cannot be deserialized
	// skipping validation

	if len(r.Items) == 0 {
		return false
	}

	for _, i := range r.Items {
		if !i.IsValid() {
			return false
		}
	}

	return validateFloat64(r.Total)
}

func validateFloat64(f float64) bool {
	// request is already parsed by the json parser
	// TODO: accurate request validation
	// for accurate validation, validate request before unmarshaling
	// once unmarshaled, precision can be lost. trailing 0 digits will be lost
	// validating the request payload is out of the scope of the domain object validation
	isMatch, err := regexp.MatchString("^\\d+\\.\\d{2}$", fmt.Sprintf("%.2f",f))
	if err != nil || !isMatch {
		return false
	}
	return true
}
