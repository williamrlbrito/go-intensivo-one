package main

import (
	"database/sql"
	"fmt"

	"github.com/williamrlbrito/go-intensivo-one/internal/order/infra/database"
	"github.com/williamrlbrito/go-intensivo-one/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	calculate := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    1,
		Price: 100,
		Tax:   10,
	}
	output, err := calculate.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The final price is: %.2f", output.TotalPrice)
}
