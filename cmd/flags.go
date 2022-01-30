package cmd

import "github.com/urfave/cli/v2"

var serverFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "region",
		Usage:    "The AWS Cognito User Pool region",
		Required: true,
	},
	&cli.StringFlag{
		Name:     "pool-id",
		Usage:    "The AWS Cognito User Pool identifier",
		Required: true,
	},
}
