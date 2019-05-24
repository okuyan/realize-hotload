package neco

import (
	"net/http"

	"github.com/labstack/echo"
)

func CallNeco(c echo.Context) error {
	return c.String(http.StatusOK, "hoge hoge neco neco")
}
