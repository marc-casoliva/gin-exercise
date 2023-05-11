package main

import (
	"encoding/json"
	"fmt"
	"gin-exercise/m/v2/domain"
	"gin-exercise/m/v2/infrastructure"
	"net/http"
	"os"

	"gin-exercise/m/v2/infrastructure/db"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var productRepository domain.ProductRepository
var asyncProducer sarama.AsyncProducer

type postReq struct {
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

func postProductHandler(ctx *gin.Context) {
	req := &postReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "malformed json"})
		return
	}

	// new kafka message
	// send message to kafka
	// iferr abort with 404 or 500

	// This goes go the kafka consumer:
	p, err := domain.NewProduct(uuid.NewString(), domain.NewPriceInEuros(req.Price), req.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	productInBytes, _ := json.Marshal(p)
	msg := sarama.ProducerMessage{
		Topic: "product",
		Key:   sarama.StringEncoder(p.ID),
		Value: sarama.ByteEncoder(productInBytes),
	}
	asyncProducer.Input() <- &msg

	ctx.JSON(http.StatusCreated, p)
}

func getProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	p, err := productRepository.RetrieveById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func initConfig() {

	viper.SetConfigFile("config/config-local.yml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	var err error
	asyncProducer, err = sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

}

func main() {
	initConfig()

	err := db.MigrateTables()
	if err != nil {
		fmt.Println("problem migrating database: ", err)
		os.Exit(1)
	}
	router := gin.Default()
	productRepository, _ = infrastructure.NewGormProductRepository()

	router.POST("/product", postProductHandler)
	router.GET("/product/:id", getProductHandler)

	router.Run("localhost:8080")
}
