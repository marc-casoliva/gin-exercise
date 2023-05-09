package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var productRepository ProductRepository

type postReq struct {
	Price       float64 `json:"price" binding:"required"`
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

func main() {
	router := gin.Default()
	productRepository = NewInMemoryProductRepository()

	router.POST("/product", postProductHandler)
	router.GET("/product/:id", getProductHandler)

	router.Run("localhost:8080")
}
