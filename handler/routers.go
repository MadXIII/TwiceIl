package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/twiceil/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	route := gin.New()

	route.GET("/products")
	prod := route.Group("/product")
	{
		// prod.GET("/add")
		prod.POST("/add")
		prod.PATCH("/edit/:productId")
	}
	cmd := route.Group("/cmd")
	{
		cmd.POST("/add-product")
		cmd.PATCH("/edit-product")
		cmd.DELETE("/delete-product")
	}
	route.POST("/q/product-search-by-name")

	return route
}
