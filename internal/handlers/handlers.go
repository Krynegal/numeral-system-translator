package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Krynegal/numeral-system-translator.git/internal/models"
	"github.com/Krynegal/numeral-system-translator.git/internal/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

type Router struct {
	*gin.Engine
}

func NewRouter(converter Converter) *Router {
	engine := gin.Default()

	engine.POST("/api/translate", TranslateHandler(converter))

	return &Router{
		engine,
	}
}

func TranslateHandler(conv Converter) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonRequest, err := c.GetRawData()
		if err != nil {
			httpError(c, http.StatusBadRequest, fmt.Errorf("cannot get request data: %w", err))

			return
		}

		var request *models.Request

		if err = json.Unmarshal(jsonRequest, &request); err != nil {
			httpError(c, http.StatusBadRequest, fmt.Errorf("cannot parse request data: %w", err))

			return
		}

		err = validators.CheckRequest(request)
		if err != nil {
			httpError(c, http.StatusBadRequest, err)

			return
		}

		res, err := conv.Convert(*request.Number, *request.Base, *request.ToBase)
		if err != nil {
			convertingError := fmt.Errorf("error while converting number %v from %v numeral system to %v: %w",
				request.Number,
				request.Base,
				request.ToBase,
				err,
			)
			httpError(c, http.StatusBadRequest, convertingError)

			return
		}

		c.JSON(http.StatusOK, models.NewResponse(res))
	}
}
