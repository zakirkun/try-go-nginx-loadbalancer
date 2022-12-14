package main

import (
	"net/http"

	app "github.com/labstack/echo/v4"
)

func main() {
	e := app.New()

	e.GET("/", func(c app.Context) error {
		return c.String(http.StatusOK, "Gateway Running!")
	})

	e.Logger.Fatal(e.Start(":5000"))
}
