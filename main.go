package main

import (
	"fmt"
	"net/http"
	"os"

	"gin-exercise/m/v2/infrastructure/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var productRepository ProductRepository

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

	p, err := NewProduct(uuid.NewString(), NewPriceInEuros(req.Price), req.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	productRepository.Save(p)

	ctx.JSON(http.StatusCreated, p)
}

func getProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	p, err := productRepository.RetreiveById(id)
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
}

func main() {
	initConfig()

	err := db.MigrateTables()
	if err != nil {
		fmt.Println("problem migrating database: ", err)
		os.Exit(1)
	}
	router := gin.Default()
	productRepository = NewInMemoryProductRepository()

	router.POST("/product", postProductHandler)
	router.GET("/product/:id", getProductHandler)

	router.Run("localhost:8080")
}
