package controller

import (
	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p productController) GetProducts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get products",
	})
}
