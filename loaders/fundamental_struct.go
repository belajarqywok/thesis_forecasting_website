package loaders

type Fundamental struct {
	FontawesomeIcon      string `json:"fontawesome_icon"`
	Symbol               string `json:"symbol"`
	SectorID             string `json:"sector_id"`
	ShortName            string `json:"shortName"`

  Address              string `json:"address"`
  Phone                string `json:"phone"`
  Website              string `json:"website"`
  MarketCap            string `json:"marketCap"`
  DividendRate         string `json:"dividendRate"`
  DividendYield        string `json:"dividendYield"`
  EarningsGrowth       string `json:"earningsGrowth"`
  ProfitMargins        string `json:"profitMargins"`
  GrossMargins         string `json:"grossMargins"`
  Beta                 string `json:"beta"`
  BookValue            string `json:"bookValue"`
  PriceToBook          string `json:"priceToBook"`

  QuickRatio           string `json:"quickRatio"`
  CurrentRatio         string `json:"currentRatio"`
  DebtToEquity         string `json:"debtToEquity"`
  RevenuePerShare      string `json:"revenuePerShare"`
  RevenueGrowth        string `json:"revenueGrowth"`
  Ebitda               string `json:"ebitda"`
  RegularMarketChange  string `json:"regularMarketChange"`
  PayoutRatio          string `json:"payoutRatio"`
  TrailingPE           string `json:"trailingPE"`
  ForwardPE            string `json:"forwardPE"`
  TrailingEps          string `json:"trailingEps"`
  ForwardEps           string `json:"forwardEps"`
}

type FundamentalsWrapper struct {
	Fundamentals Fundamental `json:"fundamentals"`
}
