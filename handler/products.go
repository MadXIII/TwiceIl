package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Products(c *gin.Context) {
	// var body model.Product
	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status": "failure",
	// 		"error":  "invalid body",
	// 	})
	// 	return
	// }
	h.service.ToGet()

	c.JSON(http.StatusOK, gin.H{
		"product": 1,
	})
}
