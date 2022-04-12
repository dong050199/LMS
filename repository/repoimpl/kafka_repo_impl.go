package repoimpl

import (
	"LMS_Kafka/model"
	"LMS_Kafka/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type KafkaRepoImpl struct {
	DB *mongo.Database
}

func NewKafkaRepo(db *mongo.Database) repository.KafkaRepo {
	return &KafkaRepoImpl{
		DB: db,
	}
}

func (config *KafkaRepoImpl) ConsumeAndProduceOrder(consumCfg model.Kafkacfg, produceCfg model.Kafkacfg, MongoCollection string) error {
	fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": consumCfg.BsServer,
		"group.id":          consumCfg.GroupId,
		"auto.offset.reset": consumCfg.AutoOsRs,
	})
	if err != nil {
		return err
	}
	c.SubscribeTopics([]string{consumCfg.Topic}, nil)
	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			data := string(msg.Value)
			var dataconsumer model.KafkaRwOrder
			err := json.Unmarshal([]byte(data), &dataconsumer)
			fmt.Println(dataconsumer)
			if err != nil {
				//panic(err)
				//return err
				fmt.Println(err)
				continue
			}

			//THIS METHOD FOR GATEWAY READ THE FIRST DOCUMENT FROM QUEUE IN SERVER.
			if dataconsumer.Method == "GET_FIRST" {
				var dataproduce model.Order
				err = config.DB.Collection(MongoCollection).FindOne(context.Background(), bson.D{{}}).Decode(&dataproduce)
				if err != nil {
					fmt.Println(err)
					continue
				}
				dataconsumer.Body = []model.Order{}
				dataconsumer.Body = append(dataconsumer.Body, dataproduce)
				fmt.Println(dataconsumer)
				//producer publish to kafka
				fmt.Println("save to kafka")
				jsonString, err := json.Marshal(dataconsumer)
				dataString := string(jsonString)
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCfg.BsServer})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// Produce messages to topic (asynchronously)
				topic := produceCfg.Topic
				for _, word := range []string{string(dataString)} {
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(word),
					}, nil)
				}
			}
			if dataconsumer.Method == "DELETE_OID" {
				var dataproduce model.Order
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				_, errr := config.DB.Collection(MongoCollection).DeleteOne(ctx, model.Order{Order_ID: dataconsumer.Id})
				if errr != nil {
					fmt.Println(err)
					continue
				}
				dataconsumer.Body = []model.Order{}
				dataconsumer.Body = append(dataconsumer.Body, dataproduce)
				fmt.Println(dataconsumer)
				//producer publish to kafka
				fmt.Println("save to kafka")
				jsonString, err := json.Marshal(dataconsumer)
				dataString := string(jsonString)
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCfg.BsServer})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// Produce messages to topic (asynchronously)
				topic := produceCfg.Topic
				for _, word := range []string{string(dataString)} {
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(word),
					}, nil)
				}
			}

			if dataconsumer.Method == "GET_ID" {
				var dataproduce model.Order
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				err := config.DB.Collection(MongoCollection).FindOne(ctx, model.Order{Order_ID: dataconsumer.Id}).Decode(&dataproduce)
				if err != nil {
					fmt.Println(err)
					continue
				}
				dataconsumer.Body = []model.Order{}
				dataconsumer.Body = append(dataconsumer.Body, dataproduce)
				fmt.Println(dataconsumer)
				//producer publish to kafka
				fmt.Println("save to kafka")
				jsonString, err := json.Marshal(dataconsumer)
				dataString := string(jsonString)
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCfg.BsServer})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// Produce messages to topic (asynchronously)
				topic := produceCfg.Topic
				for _, word := range []string{string(dataString)} {
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(word),
					}, nil)
				}
			}

			if dataconsumer.Method == "GET_ALL" {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				curso, err := config.DB.Collection(MongoCollection).Find(ctx, bson.M{})
				if err != nil {
					fmt.Println(err)
					continue
				}
				dataconsumer.Body = []model.Order{}
				for curso.Next(ctx) {
					var prodct model.Order
					curso.Decode(&prodct)
					dataconsumer.Body = append(dataconsumer.Body, prodct)

				}
				fmt.Println(dataconsumer)
				//producer publish to kafka
				fmt.Println("save to kafka")
				jsonString, err := json.Marshal(dataconsumer)
				dataString := string(jsonString)
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCfg.BsServer})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// Produce messages to topic (asynchronously)
				topic := produceCfg.Topic
				for _, word := range []string{string(dataString)} {
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(word),
					}, nil)
				}
			}

			if dataconsumer.Method == "UPDATE_OID" {
				//var dataproduce model.Order
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				_, err := config.DB.Collection(MongoCollection).UpdateOne(ctx, model.Order{Order_ID: dataconsumer.Id}, bson.M{"$set": dataconsumer.Body[0]})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// dataconsumer.Body = []model.Order{}
				// dataconsumer.Body = append(dataconsumer.Body, dataconsumer.Body[0])
				fmt.Println(dataconsumer.Body[0])
				//producer publish to kafka
				fmt.Println("save to kafka")
				jsonString, err := json.Marshal(dataconsumer)
				dataString := string(jsonString)
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCfg.BsServer})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// Produce messages to topic (asynchronously)
				topic := produceCfg.Topic
				for _, word := range []string{string(dataString)} {
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(word),
					}, nil)
				}
			}

			//THIS METHOD FOR GATEWAY DELETE FIRST DOCUMENT FROM SERVER WHEN MAKING PRODUCT DONE.
			if dataconsumer.Method == "DELETE_FIRST" {
				var dataproduce model.Order
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				_, errr := config.DB.Collection(MongoCollection).DeleteOne(ctx, bson.D{{}})
				if errr != nil {
					fmt.Println(err)
					continue
				}
				dataconsumer.Body = []model.Order{}
				dataconsumer.Body = append(dataconsumer.Body, dataproduce)
				fmt.Println(dataconsumer)
				//producer publish to kafka
				fmt.Println("save to kafka")
				jsonString, err := json.Marshal(dataconsumer)
				dataString := string(jsonString)
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCfg.BsServer})
				if err != nil {
					fmt.Println(err)
					continue
				}
				// Produce messages to topic (asynchronously)
				topic := produceCfg.Topic
				for _, word := range []string{string(dataString)} {
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(word),
					}, nil)
				}
			}

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
	c.Close()
	return nil
}

func (config *KafkaRepoImpl) ConsumeOrderAndSaveToMongo(kcfg model.Kafkacfg, MongoCollection string) error {
	fmt.Println("Start receiving from Kafka OrderPage")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kcfg.BsServer,
		"group.id":          kcfg.GroupId,
		"auto.offset.reset": kcfg.AutoOsRs,
	})
	if err != nil {
		return err
	}
	c.SubscribeTopics([]string{kcfg.Topic}, nil)
	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			data := string(msg.Value)
			var order model.Order
			err := json.Unmarshal([]byte(data), &order)
			if err != nil {
				//panic(err)
				fmt.Println(err)
				continue
			}
			//do sth here
			_, err = config.DB.Collection(MongoCollection).InsertOne(context.Background(), order)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
	c.Close()
	return nil
}
