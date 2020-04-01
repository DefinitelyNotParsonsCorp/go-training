package main

import (
	"os"

	"github.com/urfave/cli/v2"
	_ "github.com/urfave/cli/v2/altsrc"
	"go.uber.org/zap"
)

var log *zap.Logger

func setupLogging(c *cli.Context) error {

	var cfg zap.Config
	var err error

	switch c.String("log-type") {
	default:
		fallthrough
	case "prod", "production":
		cfg = zap.NewProductionConfig()
	case "dev", "development":
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.Encoding = c.String("log-encoding")

	log, err = cfg.Build()
	return err
}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "start the server",
				Action: RunServer,
				Before: setupLogging,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "listen-addr",
						Value:   ":1234",
						Usage:   "Listen on this address",
						EnvVars: []string{"LISTEN_ADDR"},
					},
					&cli.StringFlag{
						Name:    "log-type",
						Value:   "prod",
						Usage:   "Choose log type (prod, dev)",
						EnvVars: []string{"LOG_TYPE"},
					},
					&cli.StringFlag{
						Name:    "log-encoding",
						Value:   "json",
						Usage:   "Choose encoding method (json, console)",
						EnvVars: []string{"LOG_ENCODING"},
					},
				},
			},
		},
	}

	app.Run(os.Args)
	return
}
