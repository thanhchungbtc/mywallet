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
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "fresh",
					Usage: "Make fresh migration",
				},
			},
			Action: func(c *cli.Context) error {
				db, err := database.New()
				if err != nil {
					log.Fatal(err)
				}
				if c.Bool("fresh") {
					log.Print("Dropping all tables...")
					// drop all tables
					db.DropTableIfExists(model.User{})
					db.DropTableIfExists(model.Category{})
					db.DropTableIfExists(model.Account{})
					db.DropTableIfExists(model.Expense{})
					log.Print("Dropping tables done")
				}

				log.Print("Migrating...")

				db.AutoMigrate(model.User{})
				db.AutoMigrate(model.Category{})
				db.AutoMigrate(model.Account{})
				db.AutoMigrate(model.Expense{})

				log.Print("Migration done")
				return nil
			},
		},
		cli.Command{
			Name: "seed",
			Action: func(c *cli.Context) {
				//db, err := database.New()
				//if err != nil {
				//	log.Fatal(err)
				//}
				//user := model.User{
				//	Username: "admin",
				//}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
