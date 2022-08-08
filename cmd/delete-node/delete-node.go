package main

import (
	"log"
	"os"

	"github.com/hostwithquantum/clean-up-consul/pkg/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name: "Planetary Quantum GmbH",
			},
		},
		Usage: "delete all services under a `node` entity in Consul",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "consul",
				Aliases:  []string{"c"},
				Usage:    "This could be a local port-forward from Docker/Kubernetes (consul:8500)",
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "node",
				Aliases:  []string{"n"},
				Usage:    "This is what consul-sync adds itself under",
				Value:    "k8s-sync",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			utility := utils.New(cCtx.String("consul"))
			return utility.DeleteService(cCtx.String("node"), "")
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
