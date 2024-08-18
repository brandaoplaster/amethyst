package templates

func GenerationController() string {
	template := `package controllers

import "net/http"

type {{.Name}}Controller struct{}

{{range .Actions}}
func (c *{{$.Name}}Controller) {{.}}(w http.ResponseWriter, r *http.Request) {
	// Handle {{.}}
}
{{end}}
`
	return template
}
