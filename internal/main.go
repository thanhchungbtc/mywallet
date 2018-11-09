package main

import (
	"log"

	"github.com/thanhchungbtc/mywallet/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port ...")
	log.Fatal(a.Run(":3000"))
}
