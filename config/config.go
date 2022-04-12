package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
	Order       OrderConfiguration
	Gateway     GatewayConfiguration
	Mes         MesConfiguration
}

type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

type OrderConfiguration struct {
	BootStrapServer string
	GroupID         string
	AutoOffsetReset string
	TopicConsumer   string
	TopicProducer   string
	Topic           string
}

type GatewayConfiguration struct {
	BootStrapServer string
	GroupID         string
	AutoOffsetReset string
	TopicConsumer   string
	TopicProducer   string
}

type MesConfiguration struct {
	BootStrapServer string
	GroupID         string
	GroupID1        string
	AutoOffsetReset string
	TopicConsumer   string
	TopicProducer   string
}

var Cfg Configuration

func GetConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(Cfg)
}
