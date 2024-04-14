package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/tuhalang/authen/domain"
	"io"
	"net/http"
)

type ValidateController struct {
	validationUseCase domain.ValidationUseCase
}

// NewValidateController returns a new ValidationController object
func NewValidateController(validationUseCase domain.ValidationUseCase) ValidateController {
	return ValidateController{validationUseCase: validationUseCase}
}

func (vc *ValidateController) Validate(c *gin.Context) {
	log := zerolog.Ctx(c.Request.Context())

	var request domain.ValidationRequest

	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err)
		}
	}(c.Request.Body)

	validateRes, errRes := vc.validationUseCase.Validate(c.Request.Context(), request)

	if errRes != nil {
		c.JSON(errRes.HTTPCode, errRes)
		return
	}

	c.JSON(http.StatusOK, validateRes)
}
