package mtemplates

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse

	}

	return t.templates.ExecuteTemplate(w, name, data)
}
func Rend() *TemplateRenderer {
	Renderer := &TemplateRenderer{
		templates: template.Must(template.New("queue").Funcs(template.FuncMap{
			"add": add,
		}).ParseGlob("templates/*"),
		)}
	return Renderer

}
func add(x, y int) int {

	return x + y
}
