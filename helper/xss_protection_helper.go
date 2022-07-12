package helper

import (
	"html/template"

	"github.com/microcosm-cc/bluemonday"
)

func XSSProtection(inputText string) interface{} {
	parseInput := bluemonday.UGCPolicy()

	sanitizeInput := parseInput.Sanitize(inputText)

	return template.HTMLEscapeString(sanitizeInput)
}
