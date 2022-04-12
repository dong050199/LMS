package repository

import (
	"LMS_Kafka/model"
)

type KafkaRepo interface {
	ConsumeOrderAndSaveToMongo(kcfg model.Kafkacfg, MongoCollection string) error
	ConsumeAndProduceOrder(consumCfg model.Kafkacfg, produceCfg model.Kafkacfg, MongoCollection string) error
}
