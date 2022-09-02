package router

import (
	controller "fita-test/controller/product"
	"fita-test/service/product"

	"github.com/gin-gonic/gin"
)

func Route(v1 *gin.RouterGroup, prd *product.Service) {
	handler := controller.ControllerHandler(prd)
	product := v1.Group("/product")
	{
		product.POST("/add-to-cart", handler.AddToCart)
		product.POST("/add", handler.Add)
		product.POST("/checkout", handler.Checkout)
	}
}