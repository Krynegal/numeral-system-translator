package handlers

import (
	"github.com/Krynegal/numeral-system-translator.git/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	engine := gin.Default()

	engine.POST("/api/translate", translateHandler)

	return &Router{
		engine,
	}
}

func translateHandler(c *gin.Context) {
	const defaultValue = "0"

	response := models.NewResponse(defaultValue)

	c.JSON(http.StatusOK, response)
}
