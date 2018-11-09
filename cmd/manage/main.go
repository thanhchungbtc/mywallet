package main

import (
	"os"

	"github.com/thanhchungbtc/mywallet/internal/model"

	"github.com/thanhchungbtc/mywallet/internal/database"

	"github.com/labstack/gommon/log"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "CLI for mywallet app"

	app.Commands = cli.Commands{
		cli.Command{
			Name:    "migrate",
			Aliases: []string{"m"},
			Action: func(c *cli.Context) error {
				db, err := database.New()
				if err != nil {
					log.Fatal(err)
				}
				db.AutoMigrate(model.User{})
				db.AutoMigrate(model.Category{})
				db.AutoMigrate(model.Account{})
				db.AutoMigrate(model.Expense{})
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
