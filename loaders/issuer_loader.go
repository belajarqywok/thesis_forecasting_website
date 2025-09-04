package loaders

import (
	"os"
	"encoding/json"
)

type Issuer struct {
	FontawesomeIcon string `json:"fontawesome_icon"`
	Symbol          string `json:"symbol"`
	SectorID        string `json:"sector_id"`
	ShortName       string `json:"shortName"`
	Beta            string `json:"beta"`
	DividendYield   string `json:"dividendYield"`
}

type IssuersWrapper struct {
	Issuers []Issuer `json:"infographics"`
}

func IssuerLoader(path string) ([]Issuer, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data IssuersWrapper
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Issuers, nil
}
