package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"lodestar/config"
	"net/http"
	"text/template"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/urfave/cli/v2"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func commandServer() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Run server",
		Flags:   serverFlags,
		Action: func(c *cli.Context) error {
			// Initialize config
			config.Initialize(c)

			// Launch server
			e := echo.New()
			e.Use(middleware.Logger())
			e.Use(middleware.Recover())
			e.Use(middleware.Secure())
			e.Use(middleware.Gzip())

			t := &Template{
				templates: template.Must(template.ParseGlob("public/views/*.html")),
			}
			e.Renderer = t

			e.GET("/", func(c echo.Context) error {
				return c.Render(http.StatusOK, "index.html", "Lodestar")
			})

			api := e.Group("/api")
			api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
				KeyFunc: getKey,
			}))
			api.GET("/", func(c echo.Context) error {
				token := c.Get("user")
				fmt.Printf("Context: %#v\n", token)
				return c.String(http.StatusOK, "Hello, Fam!")
			})

			e.Logger.Fatal(e.Start(":8082"))

			return nil
		},
	}
}

func getKey(token *jwt.Token) (interface{}, error) {
	url := fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json",
		config.AppConfig.Region,
		config.AppConfig.PoolID,
	)

	// TODO: No need to fetch the JWK for every request
	keySet, err := jwk.Fetch(context.Background(), url)
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}
