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

	route.GET("/products", h.Products)

	cmd := route.Group("/cmd")
	{
		cmd.POST("/add-product", h.Create)
		cmd.PUT("/edit-product", h.Update)
		cmd.DELETE("/delete-product", h.Delete)
	}
	route.POST("/q/product-search-by-name", h.Find)

	return route
}
