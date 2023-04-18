package main

import (
	"GO-Payment/internal/delivery"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	router := delivery.NewAppServer()

	if err := router.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
