package controller

import (
	"LMS_Kafka/config"
	"LMS_Kafka/driver"
	"LMS_Kafka/model"
	"LMS_Kafka/repository/repoimpl"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PublishToKafka(kafkatopic string, bsServer string, data model.KafkaRwOrder) error {
	fmt.Println("Producer to kafka topic")
	jsonString, err := json.Marshal(data)
	dataString := string(jsonString)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bsServer})
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Produce messages to topic (asynchronously)
	topic := kafkatopic
	for _, word := range []string{string(dataString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	return nil
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	ProductRepo := repoimpl.NewOrderRepo(driver.Mongo.Client.Database(config.Cfg.Mongo.Database))
	data, err := ProductRepo.SelectId(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	ProductRepo := repoimpl.NewOrderRepo(driver.Mongo.Client.Database(config.Cfg.Mongo.Database))
	data, err := ProductRepo.Select()
	fmt.Println(data)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var cfg_kafka model.KafkaRwOrder
	cfg_kafka.Method = "DELETE_OID"
	cfg_kafka.Id = id
	err := PublishToKafka(config.Cfg.Mes.TopicConsumer, config.Cfg.Mes.BootStrapServer, cfg_kafka)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	//id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	//ProductRepo := repoimpl.NewOrderRepo(driver.Mongo.Client.Database("DONG_TEST"))
	//err := ProductRepo.Delete(id)
	json.NewEncoder(w).Encode("delete_done")
}

// func CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "applycation/json")
// 	var product models.Order
// 	_ = json.NewDecoder(r.Body).Decode(&product)
// 	ProductRepo := repoimpl.NewOrderRepo(driver.Mongo.Client.Database("DONG_TEST"))
// 	err := ProductRepo.Insert(product)
// 	fmt.Println(product)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(product)
// }

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var cfg_kafka model.KafkaRwOrder
	cfg_kafka.Method = "UPDATE_OID"
	cfg_kafka.Id = id
	err := PublishToKafka(config.Cfg.Mes.TopicConsumer, config.Cfg.Mes.BootStrapServer, cfg_kafka)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode("update_done")

}
