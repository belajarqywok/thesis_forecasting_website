package loaders

type Historical struct {
	Date           string   `json:"date"`
	FullDate       string   `json:"full_date"`
	Open          float64   `json:"open"`
	High          float64   `json:"high"`
	Low           float64   `json:"low"`
	Close         float64   `json:"close"`
  Volume        float64   `json:"volume"`
}

type HistoricalsWrapper struct {
	Historicals   []Historical   `json:"historicals"`
}
