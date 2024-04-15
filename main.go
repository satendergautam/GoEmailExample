package main

import (
	"log"
	"main/mappings"

	gotenv "github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mappings.CreateUrlMappings()
	err := mappings.Router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
