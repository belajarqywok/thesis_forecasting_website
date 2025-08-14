package handlers

import (
  "os"
  "fmt"
  "sync"
  "github.com/gofiber/fiber/v2"
	
  loaders   "thesis_forecasting_website/loaders"
)

func InfographicHandler(c *fiber.Ctx) error {
	issuer_name := c.Query("issuer_name")
  if issuer_name == "" {
    return c.Redirect("/")
  }

  // json validation
  fundamental_json := fmt.Sprintf("./indonesia_stocks/fundamentals/%s.json", issuer_name)
  historical_json  := fmt.Sprintf("./indonesia_stocks/historicals/%s.json",  issuer_name)
  indicator_json   := fmt.Sprintf("./indonesia_stocks/indicators/%s.json",   issuer_name)

  json_paths  := []string{fundamental_json, historical_json, indicator_json}
  err_channel := make(chan error, len(json_paths))

  var waitgroup_validate sync.WaitGroup
  waitgroup_validate.Add(len(json_paths))

  for _, json_path := range json_paths {
    go func(path string) {
      defer waitgroup_validate.Done()
      if _, err := os.Stat(path); os.IsNotExist(err) {
        err_channel <- fmt.Errorf("file missing: %s", path)
        return
      }
      err_channel <- nil
    }(json_path)
  }

  waitgroup_validate.Wait()

  for i := 0; i < len(json_paths); i++ {
    if err := <-err_channel; err != nil {
      return c.Redirect("/")
    }
  }

  // load infographics (fundamental, historical, indicator)
  var (
    fundamental   *loaders.Fundamental
    historicals   []loaders.Historical
    indicators    []loaders.Indicator
    errors        []error
    mutex         sync.Mutex
  )

  addError := func(err error) {
    if err != nil {
      mutex.Lock()
      errors = append(errors, err)
      mutex.Unlock()
    }
  }

  var waitgroup_infographics sync.WaitGroup
  waitgroup_infographics.Add(3)

  go func() {
    defer waitgroup_infographics.Done()
    data, err := loaders.FundamentalLoader(fundamental_json)
    if err != nil { addError(err) }

    mutex.Lock()
    fundamental = data
    mutex.Unlock()
  }()

  go func() {
    defer waitgroup_infographics.Done()
    data, err := loaders.HistoricalLoader(historical_json)
    if err != nil { addError(err) }

    mutex.Lock()
    historicals = data
    mutex.Unlock()
  }()

  go func() {
    defer waitgroup_infographics.Done()
    data, err := loaders.IndicatorLoader(indicator_json)
    if err != nil { addError(err) }

    mutex.Lock()
    indicators = data
    mutex.Unlock()
  }()

  waitgroup_infographics.Wait()

  if len(errors) > 0 {
    for _, err := range errors { println(err.Error()) }
    return c.Redirect("/")
  }

  data := fiber.Map{
    "Fundamentals": fundamental,
    "Historicals":  historicals,
    "Technicals":   indicators,
  }

  return c.Render("infographic", data)
}
