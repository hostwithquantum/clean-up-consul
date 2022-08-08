package main

import (
	"fmt"
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
		Usage: "delete all services by name or by tag from Consul",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "consul",
				Aliases:  []string{"c"},
				Usage:    "This could be a local port-forward from Docker/Kubernetes (consul:8500)",
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "service",
				Aliases: []string{"s"},
				Usage:   "E.g. `haproxy-ingress-metrics-ingress-controller`",
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "tag",
				Aliases: []string{"t"},
				Value:   "k8s",
			},
		},
		Action: func(cCtx *cli.Context) error {
			var deleteSvc []string

			utility := utils.New(cCtx.String("consul"))

			if cCtx.String("service") == "" {
				deleteSvc = utility.FindServicesToDelete(cCtx.String("tag"))
				if len(deleteSvc) == 0 {
					fmt.Println("Nothing to do. Bye.")
					return nil
				}

				fmt.Println("About to delete some services...")
			} else {
				deleteSvc = append(deleteSvc, cCtx.String("service"))
			}

			for _, svc := range deleteSvc {
				fmt.Printf("Will delete: %s\n", svc)

				fmt.Println("Getting service IDs!")

				aSvc := utility.GetService(svc)

				for _, actualSvc := range aSvc {
					serviceId := actualSvc["ServiceID"].(string)
					node := actualSvc["Node"].(string)
					fmt.Printf("Found a serviceID: %s\n", serviceId)
					fmt.Printf("Found a node: %s\n", node)
					fmt.Println("")

					err := utility.DeleteService(node, serviceId)
					if err != nil {
						return err
					}
				}

				// id := aSvc["ID"].(string)
				// fmt.Printf("Found service ID: %s", id)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
