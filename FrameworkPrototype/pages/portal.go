//go:build amd64 && !wasm
// +build amd64,!wasm

package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"

	"third"
)

var number int

func PortalPage(c echo.Context) error {
	number++
	return c.Render(http.StatusOK, "portal", template.HTML(fmt.Sprintf("<p>[SERVER SIDE] %s %d</p>", third.Third("SERVER"), number)))
}
