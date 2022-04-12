package repoimpl

import (
	"LMS_Kafka/config"
	models "LMS_Kafka/model"
	repo "LMS_Kafka/repository"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepoImpl struct {
	DB *mongo.Database
}

func NewOrderRepo(db *mongo.Database) repo.OrderRepo {
	return &OrderRepoImpl{
		DB: db,
	}
}

func (mongo *OrderRepoImpl) Update(u models.Order, id primitive.ObjectID) error {
	var product models.Order
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.DB.Collection(config.Cfg.Mongo.Collection).UpdateOne(ctx, models.Order{ID: id}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

func (mongo *OrderRepoImpl) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.DB.Collection(config.Cfg.Mongo.Collection).DeleteOne(ctx, models.Order{ID: id})
	if err != nil {
		return err
	}
	return nil
}

func (mongo *OrderRepoImpl) Select() ([]models.Order, error) {
	var products []models.Order
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	curso, err := mongo.DB.Collection(config.Cfg.Mongo.Collection).Find(ctx, bson.D{{}})
	if err != nil {
		return products, err
	}
	for curso.Next(ctx) {
		var prodct models.Order
		fmt.Println(curso)
		curso.Decode(&prodct)
		fmt.Println(prodct)
		products = append(products, prodct)

	}
	return products, nil
}

func (mongo *OrderRepoImpl) SelectId(id primitive.ObjectID) (models.Order, error) {
	var product models.Order
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := mongo.DB.Collection(config.Cfg.Mongo.Collection).FindOne(ctx, models.Order{ID: id}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (mongo *OrderRepoImpl) Insert(u models.Order) error {
	bbytes, _ := bson.Marshal(u)
	_, err := mongo.DB.Collection(config.Cfg.Mongo.Collection).InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}
	return nil
}
