package loaders

import (
	"os"
	"encoding/json"
)

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
