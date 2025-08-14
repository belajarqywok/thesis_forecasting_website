package loaders

type Indicator struct {
	Date            string  `json:"date"`
	FullDate        string  `json:"full_date"`
	MFI            float64  `json:"MFI"`
	RSI            float64  `json:"RSI"`
	MACD           float64  `json:"MACD"`
}

type IndicatorsWrapper struct {
	Indicators   []Indicator   `json:"indicators"`
}
