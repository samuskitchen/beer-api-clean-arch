package domain

// CurrencyLayer struct
type CurrencyLayer struct {
	Success   bool                   `json:"success"`
	Terms     string                 `json:"terms"`
	Privacy   string                 `json:"privacy"`
	Timestamp int                    `json:"timestamp"`
	Source    string                 `json:"source"`
	Quotes    map[string]interface{} `json:"quotes"`
}

// CurrencyLayerRepository interface repository
type CurrencyLayerRepository interface {
	GetCurrency(currency, currencyBeer string) ([]float64, error)
}
