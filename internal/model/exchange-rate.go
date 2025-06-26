package model

type ExchangeRate struct {
	ID             int
	BaseCurrency   Currency `json:"baseCurrency"`
	TargetCurrency Currency `json:"targetCurrency"`
	Rate           float32  `json:"rate"`
}
