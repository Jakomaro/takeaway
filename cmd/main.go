package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jakomaro/takeaway/internal/repository"
	"github.com/jakomaro/takeaway/internal/services"
)

func main() {

	connString := "postgresql://postgres:postgres@localhost:5432/takeaway?sslmode=disable"
	postgresDB, err := repository.NewPostgresDB(connString)
	if err != nil {
		log.Println(err)
	}

	ESClient, err := repository.NewElasticDB()
	if err != nil {
		log.Println(err)
	}

	var servizi []services.MenuServicer
	servizi = append(servizi, services.NewPGMenuService(postgresDB))
	servizi = append(servizi, services.NewESMenuService(ESClient))

	for _, s := range servizi {
		menu, err := s.GetMenu(context.Background())
		if err != nil {
			log.Println(err)
		}
		fmt.Println(menu)
	}
}
