package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thanhchungbtc/mywallet/server/app/database"

	_ "github.com/lib/pq"
	"github.com/thanhchungbtc/mywallet/server/app"
)

func main() {
	DB, err := database.New("postgres", "postgres://thanhchungbui@localhost/mywallet?sslmode=disable")
	a, err := app.New(DB)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf(fmt.Sprintf("Listening on port %s ...", port))
	log.Fatal(a.Run(fmt.Sprintf(":%s", port)))
}
