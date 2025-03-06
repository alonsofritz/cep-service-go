package entity

import "time"

type Address struct {
	CEP       string    `json:"cep" bson:"cep"`
	Street    string    `json:"logradouro" bson:"street"`
	City      string    `json:"localidade" bson:"city"`
	State     string    `json:"uf" bson:"state"`
	Provider  string    `json:"provider" bson:"provider"`
	FetchedAt time.Time `bson:"fetched_at"`
}
