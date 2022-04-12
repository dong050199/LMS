package controller

import (
	"LMS_Kafka/config"
	"LMS_Kafka/driver"
	"LMS_Kafka/model"
	"LMS_Kafka/repository/repoimpl"
	"fmt"
)

func Order_Consume() {
	go func() {
		var kafkacfg model.Kafkacfg
		kafkacfg.Topic = config.Cfg.Order.Topic
		kafkacfg.AutoOsRs = config.Cfg.Order.AutoOffsetReset
		kafkacfg.GroupId = config.Cfg.Order.GroupID
		kafkacfg.BsServer = config.Cfg.Order.BootStrapServer
		newkafka := repoimpl.NewKafkaRepo(driver.Mongo.Client.Database(config.Cfg.Mongo.Database))
		fmt.Println(newkafka.ConsumeOrderAndSaveToMongo(kafkacfg, config.Cfg.Mongo.Collection))
	}()
}

func Order_Request_Response() {
	go func() {
		var consum model.Kafkacfg
		consum.Topic = config.Cfg.Order.TopicConsumer
		consum.AutoOsRs = config.Cfg.Order.AutoOffsetReset
		consum.GroupId = config.Cfg.Order.GroupID
		consum.BsServer = config.Cfg.Order.BootStrapServer

		var produce model.Kafkacfg
		produce.Topic = config.Cfg.Order.TopicProducer
		produce.AutoOsRs = config.Cfg.Order.AutoOffsetReset
		produce.GroupId = config.Cfg.Order.GroupID
		produce.BsServer = config.Cfg.Order.BootStrapServer
		newkafka := repoimpl.NewKafkaRepo(driver.Mongo.Client.Database(config.Cfg.Mongo.Database))
		fmt.Println(newkafka.ConsumeAndProduceOrder(consum, produce, config.Cfg.Mongo.Collection))
	}()
}

func Gateway_Request_Response() {
	go func() {
		var consum model.Kafkacfg
		consum.Topic = config.Cfg.Gateway.TopicConsumer
		consum.AutoOsRs = config.Cfg.Gateway.AutoOffsetReset
		consum.GroupId = config.Cfg.Gateway.GroupID
		consum.BsServer = config.Cfg.Gateway.BootStrapServer

		var produce model.Kafkacfg
		produce.Topic = config.Cfg.Gateway.TopicProducer
		produce.AutoOsRs = config.Cfg.Gateway.AutoOffsetReset
		produce.GroupId = config.Cfg.Gateway.GroupID
		produce.BsServer = config.Cfg.Gateway.BootStrapServer
		newkafka := repoimpl.NewKafkaRepo(driver.Mongo.Client.Database(config.Cfg.Mongo.Database))
		fmt.Println(newkafka.ConsumeAndProduceOrder(consum, produce, config.Cfg.Mongo.Collection))
	}()
}

func MES_Request_Response() {
	go func() {
		var consum model.Kafkacfg
		consum.Topic = config.Cfg.Mes.TopicConsumer
		consum.AutoOsRs = config.Cfg.Mes.AutoOffsetReset
		consum.GroupId = config.Cfg.Mes.GroupID
		consum.BsServer = config.Cfg.Mes.BootStrapServer

		var produce model.Kafkacfg
		produce.Topic = config.Cfg.Mes.TopicProducer
		produce.AutoOsRs = config.Cfg.Mes.AutoOffsetReset
		produce.GroupId = config.Cfg.Mes.GroupID
		produce.BsServer = config.Cfg.Mes.BootStrapServer
		newkafka := repoimpl.NewKafkaRepo(driver.Mongo.Client.Database(config.Cfg.Mongo.Database))
		fmt.Println(newkafka.ConsumeAndProduceOrder(consum, produce, config.Cfg.Mongo.Collection))
	}()
}
