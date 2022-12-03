package utils

import (
	"strings"
	"text/template"
)

// Format helps with string concatenation. Example: utils.Format("Variable {{.}} ", data)
func Format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	return b.String()
}
