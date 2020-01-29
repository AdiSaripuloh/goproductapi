package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AdiSaripuloh/goproductapi/common/resolvers"
	productHandlers "github.com/AdiSaripuloh/goproductapi/modules/product/handlers"

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
	port := os.Getenv("PORT")
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	router.LoadHTMLGlob("public/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	api := router.Group("api")
	{
		api.GET("/product/:id", productHandler.GetByUUID)
		api.POST("/product", productHandler.Submit)
		api.GET("/products", productHandler.GetAll)
	}

	if err := router.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
