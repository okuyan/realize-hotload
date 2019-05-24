package main

import (
	"net/http"

	"github.com/labstack/echo"

	"app/neco"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/neco", neco.CallNeco)
	e.Logger.Fatal(e.Start(":1323"))
}
