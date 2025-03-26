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

	var sliceOfServices []services.MenuServicer
	sliceOfServices = append(sliceOfServices, services.NewSMenuService())
	sliceOfServices = append(sliceOfServices, services.NewPGMenuService(postgresDB))
	sliceOfServices = append(sliceOfServices, services.NewESMenuService(ESClient))

	for _, s := range sliceOfServices {
		fmt.Printf("%T\t\t: ", s)
		menu, err := s.GetMenu(context.Background())
		if err != nil {
			log.Println(err)
		}
		fmt.Println(menu)
	}
}
