package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func New() {
	e := echo.New()

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("wasm/views/*.html")),
	}

	e.GET("/", Index)
	e.Static("/assets", "assets")
	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
