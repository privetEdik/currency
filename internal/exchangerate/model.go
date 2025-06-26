package currency

import "currency/internal/currency"

type ExchangeRate struct {
	ID             int
	BaseCurrency   currency.Currency `json:"baseCurrency"`
	TargetCurrency currency.Currency `json:"targetCurrency"`
	Rate           float32           `json:"rate"`
}
