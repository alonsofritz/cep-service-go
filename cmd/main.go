package main

import (
	"fmt"
	"net/http"

	"github.com/alonsofritz/cep-service-go/config"
	"github.com/alonsofritz/cep-service-go/internal/adapter"
	"github.com/alonsofritz/cep-service-go/internal/repository"
	"github.com/alonsofritz/cep-service-go/internal/usecase"
)

func main() {
	db := config.ConnectDB()
	repo := repository.NewAddressRepository(db)
	uc := usecase.NewAddressUseCase(repo)
	handler := adapter.NewAddressHandler(uc)

	http.HandleFunc("/cep", handler.GetAddressHandler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
