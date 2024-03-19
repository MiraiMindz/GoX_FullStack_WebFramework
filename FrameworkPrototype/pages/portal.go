//go:build amd64 && !wasm
// +build amd64,!wasm

package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

var number int

func PortalPage(c echo.Context) error {
	number++
	return c.Render(http.StatusOK, "portal", template.HTML(fmt.Sprintf("<p>this is being rendered on the server %d</p>", number)))
}
