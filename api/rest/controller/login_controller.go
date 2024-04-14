package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/tuhalang/authen/domain"
	"io"
	"net/http"
)

// LoginController is a loginController object
type LoginController struct {
	loginUseCase domain.LoginUseCase
}

// NewLoginController init a login controller
func NewLoginController(loginUseCase domain.LoginUseCase) LoginController {
	return LoginController{
		loginUseCase: loginUseCase,
	}
}

// Login handle login flow
func (lc *LoginController) Login(c *gin.Context) {

	log := zerolog.Ctx(c.Request.Context())

	var request domain.LoginRequest

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

	log.Info().Msgf("Handle login for user %s", request.Username)

	loginRes, errRes := lc.loginUseCase.LoginByPassword(c.Request.Context(), request.Username, request.Password)

	if errRes != nil {
		c.JSON(errRes.HTTPCode, errRes)
		return
	}

	c.JSON(http.StatusOK, loginRes)
}
