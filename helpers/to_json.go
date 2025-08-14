package helpers

import (
	"encoding/json"
	"html/template"
)

func ToJSON(v interface{}) template.JS {
	b, _ := json.Marshal(v)
	return template.JS(b)
}
