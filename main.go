package main

import (
	"LMS_Kafka/config"
	"LMS_Kafka/controller"
	"LMS_Kafka/driver"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.GetConfig()
	fmt.Println(config.Cfg.Mongo.Collection)
	fmt.Println("cocl")
	//connect mongoDB onetime
	driver.ConnectMongoDB()
	//run kafka controller here
	fmt.Println("System Started Please Wait a sencond for kafka connect to server .....10s")
	controller.Order_Consume()
	controller.Order_Request_Response()
	controller.MES_Request_Response()
	controller.Gateway_Request_Response()
	// go func() {
	// 	for {
	// 		fmt.Println("hello cu")
	// 		var data model.KafkaRwOrder
	// 		data.Method = "GET_ALL"
	// 		jsonString, err := json.Marshal(data)
	// 		dataString := string(jsonString)
	// 		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.GATEWAY_BOOTSTRAP_SERVER})
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			continue
	// 		}
	// 		// Produce messages to topic (asynchronously)
	// 		topic := config.GATEWAY_TOPIC_CONSUMER
	// 		for _, word := range []string{string(dataString)} {
	// 			p.Produce(&kafka.Message{
	// 				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 				Value:          []byte(word),
	// 			}, nil)
	// 		}
	// 		time.Sleep(5 * time.Second)
	// 	}
	// }()

	//run mux api router for config
	r := mux.NewRouter()
	r.HandleFunc("/get/{id}", controller.GetOrderById).Methods("GET")
	r.HandleFunc("/getall", controller.GetAllOrder).Methods("GET")
	r.HandleFunc("/delete/{id}", controller.DeleteOrder).Methods("DELETE")
	//r.HandleFunc("/post", controller.CreateOrder).Methods("POST")
	r.HandleFunc("/put/{id}", controller.UpdateOrder).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8101", r))
}
