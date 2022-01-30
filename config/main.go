package config

import "github.com/urfave/cli/v2"

// AppConfig is a singleton application configuration
var AppConfig Configuration

// Configuration captures the app config details
type Configuration struct {
	Region string
	PoolID string
}

// Initialize creates the application configuration
func Initialize(ctx *cli.Context) {
	AppConfig = Configuration{
		Region: ctx.String("region"),
		PoolID: ctx.String("pool-id"),
	}
}
