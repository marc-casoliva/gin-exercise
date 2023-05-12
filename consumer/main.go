package main

import (
	"encoding/json"
	"fmt"
	"gin-exercise/m/v2/domain"
	"gin-exercise/m/v2/infrastructure"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func initConfig() {

	viper.SetConfigFile("config/config-local.yml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
func main() {
	initConfig()
	productRepository, _ := infrastructure.NewGormProductRepository()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	topic := "product"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	notKilled := true

	for notKilled {
		// blocks before select until there is a msg in at least one chan
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d: '%s': '%s'\n", msg.Offset, string(msg.Key), string(msg.Value))
			var p domain.Product
			if err := json.Unmarshal(msg.Value, &p); err != nil {
				log.Printf("Unmarshal for %v failed with error: %v", msg.Key, err)
			}
			productRepository.Save(p)
			consumed++
		case <-signals:
			notKilled = false
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
