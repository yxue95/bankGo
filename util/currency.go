package util

const (
	// USD is the currency code of US Dollar
	USD = "USD"
	// EUR is the currency code of Euro
	EUR = "EUR"
	// CAD is the currency code of Canadian Dollar
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
