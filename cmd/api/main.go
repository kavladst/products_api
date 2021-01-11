package main

import (
	"log"

	"github.com/kavladst/products_api/internal/app/api"
)

func main() {
	apiApplication, err := api.NewApi()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal(apiApplication.Run())
	}
}
