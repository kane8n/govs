package main

import (
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "govs",
		Usage:  "usage",
		Flags:  []cli.Flag{},
		Action: switchVersion,
		Commands: []*cli.Command{
			{
				Name:   "install",
				Usage:  "install new version",
				Action: install,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createGovsDir(home string) error {
	if err := os.MkdirAll(path.Join(home, ".govs"), 0755); err != nil {
		return err
	}
	return nil
}

func getSDKDirPath(home string) string {
	return path.Join(home, "sdk")
}
