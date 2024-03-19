//go:build amd64 && !wasm
// +build amd64,!wasm

package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PortalPage(c echo.Context) error {
	return c.Render(http.StatusOK, "portal", template.HTML("<p>this is being rendered on the server</p>"))
}
