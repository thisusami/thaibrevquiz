package repositories

import (
	"context"
	"log"

	"github.com/thisusami/thaibrevquiz/db"
	"github.com/thisusami/thaibrevquiz/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository struct {
	MongoDbProperty *db.MongoDbProperty
}

func (r *Repository) Insert(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.MongoDbProperty.Timeout)
	defer cancel()
	log.Printf("{action:REQUEST, route:insert,message:%v}", user)
	collection := r.MongoDbProperty.Client.Database(r.MongoDbProperty.DB).Collection("users")
	_, err := collection.InsertOne(ctx, bson.M{
		"username": user.Username,
		"password": user.Password,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("{action:RESPONSE, route:insert}")
	return user, nil
}
func (r *Repository) Get(filter any) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), r.MongoDbProperty.Timeout)
	defer cancel()
	log.Printf("{action:REQUEST, route:get,message:%v}", filter)
	collection := r.MongoDbProperty.Client.Database(r.MongoDbProperty.DB).Collection("users")
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	log.Printf("{action:RESPONSE, route:get,message:%v}", user)
	return &user, nil

}
func NewRepository(mongoDbProperty *db.MongoDbProperty) *Repository {
	return &Repository{
		MongoDbProperty: mongoDbProperty,
	}
}
