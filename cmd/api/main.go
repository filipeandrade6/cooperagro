package main

import (
	"fmt"
	"log"

	"github.com/filipeandrade6/cooperagro/domain/usecases/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecases/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecases/product"
	"github.com/filipeandrade6/cooperagro/domain/usecases/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/infrastructure/repository/postgres"
)

func main() {
	dataSourceName := fmt.Sprintf("url de conexao")
	db, err := postgres.NewPostgresRepo(dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	baseProductService := baseproduct.NewService(db)
	inventoryService := inventory.NewService(db)
	productService := product.NewService(db)
	unitOfMeasureService := unitofmeasure.NewService(db)
	userService := user.NewService(db)

	middleware
	handlers

}
