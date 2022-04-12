package repository

import (
	models "LMS_Kafka/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRepo interface {
	SelectId(id primitive.ObjectID) (models.Order, error)
	Select() ([]models.Order, error)
	Insert(u models.Order) error
	Update(u models.Order, id primitive.ObjectID) error
	Delete(id primitive.ObjectID) error
}
