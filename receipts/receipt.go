package receipts

import (
	"regexp"
	"strconv"

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

	isMatch, err = regexp.MatchString("^\\d+\\.\\d{2}$", strconv.FormatFloat(i.Price, 'f', 3, 64))
	if err != nil || !isMatch {
		return false
	}

	return true
}

func (r Receipt) IsValid() bool {
	if r.Retailer == "" {
		return false
	}

	isMatch, err := regexp.MatchString("^\\S+$", r.Retailer)
	if err != nil || !isMatch {
		return false
	}

	// invalid date and time formats cannot be deserialized
	// skipping validation

	if len(r.Items) == 0 {
		return false
	}

	// total is already parsed by the json parser
	// TODO: precision validation
	// for accurate validation, validate request before unmarshaling
	// formating to 3 decimal places to handle rounding errors
	isMatch, err = regexp.MatchString("^\\d+\\.\\d{2}$", strconv.FormatFloat(r.Total, 'f', 3, 64))
	if err != nil || !isMatch {
		return false
	}

	for _, i := range r.Items {
		if !i.IsValid(){
			return false
		}
	}

	return true
}
