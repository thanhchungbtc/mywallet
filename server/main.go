package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thanhchungbtc/mywallet/server/app"
)

func main() {
	a, err := app.New()
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
