package cmd

import (
	"context"
	"errors"
	"fmt"
	"lodestar/config"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/urfave/cli/v2"
)

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
			e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
				KeyFunc: getKey,
			}))

			e.GET("/", func(c echo.Context) error {
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
