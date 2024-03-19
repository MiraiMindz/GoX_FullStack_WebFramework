//go:build amd64 && !wasm
// +build amd64,!wasm

package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Println("Test 3")
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Renderer = t

	e.File("/", "public/index.html")
	e.Static("/static", "assets")

	e.GET("/portal", PortalPage)

	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
