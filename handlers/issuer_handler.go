package handlers

import (
  "sync"
  "github.com/gofiber/fiber/v2"
	
  loaders "thesis_forecasting_website/loaders"
)

func IssuerHandler(c *fiber.Ctx) error {
	var (
		issuers      []loaders.Issuer
		sectors      []string

		mutex        sync.Mutex
		waitgroup    sync.WaitGroup
		errors       []error
	)

	addError := func(err error) {
		if err != nil {
			mutex.Lock()
			errors = append(errors, err)
			mutex.Unlock()
		}
	}

	waitgroup.Add(2)

	// issuer goroutine
	go func() {
		defer waitgroup.Done()
		issuer_data, err := loaders.IssuerLoader("./indonesia_stocks/top_50_stocks.json")
		if err != nil {
			addError(err)
			issuer_data = []loaders.Issuer{}
		}
		mutex.Lock()
		issuers = issuer_data
		mutex.Unlock()
	}()

	// sector goroutine
	go func() {
		defer waitgroup.Done()
		sector_data, err := loaders.SectorLoader("./indonesia_stocks/sectors.json")
		if err != nil {
				addError(err)
				sector_data = []string{}
		}
		mutex.Lock()
		sectors = sector_data
		mutex.Unlock()
	}()

	waitgroup.Wait()

	if len(errors) > 0 {
		for _, err := range errors {
			println(err.Error())
		}
	}

	data := fiber.Map{
		"Sectors":      sectors,
		"Infographics": issuers,
	}

	return c.Render("issuer", data)
}
