package main

import (
	"net/http"

	"github.com/goproductapi/common/resolvers"
	productHandlers "github.com/goproductapi/modules/product/handlers"

	"github.com/gin-gonic/gin"
)

var (
	productHandler *productHandlers.ProductHandler
)

func init() {
	resolver := resolvers.NewProductResolver()

	productHandler = productHandlers.NewProductHandler(resolver)
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Hello world!",
		})
	})

	api := router.Group("api")
	{
		api.GET("/product/:id", productHandler.GetByUUID)
		api.POST("/product", productHandler.Submit)
		api.GET("/products", productHandler.GetAll)
	}

	router.Run(":8000")
}
