package repository

import (
	"context"
	"time"

	"github.com/alonsofritz/cep-service-go/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressRepository struct {
	db *mongo.Collection
}

func NewAddressRepository(db *mongo.Collection) *AddressRepository {
	return &AddressRepository{db: db}
}

func (r *AddressRepository) FetchFromDB(cep string) (*entity.Address, error) {
	var addr entity.Address
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"cep": cep}
	err := r.db.FindOne(ctx, filter).Decode(&addr)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func (r *AddressRepository) SaveToDB(addr entity.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.db.InsertOne(ctx, addr)
	return err
}
