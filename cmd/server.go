package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v2"
)

func commandServer() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Run server",
		Action: func(c *cli.Context) error {
			e := echo.New()
			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "Hello, Fam!")
			})
			e.Logger.Fatal(e.Start(":8082"))

			return nil
		},
	}
}
