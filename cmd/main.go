package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/g6123/jtd-tools/pkg"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "jtd-tools",
		Commands: []*cli.Command{
			{
				Name:      "schema",
				Usage:     "Converts JSON typedef to JSON schema",
				ArgsUsage: "<file>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Writes converted JSON schema to `FILE`",
					},
				},
				Action: func(ctx *cli.Context) error {
					doc, err := pkg.ParseFile(ctx.Args().First())
					if err != nil {
						return err
					}

					schema := pkg.ToSchema(*doc)
					output, err := json.MarshalIndent(schema, "", "  ")
					if err != nil {
						return err
					}

					output_file := ctx.String("output")
					if output_file != "" {
						err = os.WriteFile(output_file, output, os.FileMode(0o644))
					} else {
						_, err = os.Stdout.Write(output)
					}

					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
