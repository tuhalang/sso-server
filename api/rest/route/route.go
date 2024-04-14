package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tuhalang/authen/api/rest/controller"
	"github.com/tuhalang/authen/internal/logger"
)

// RestRoute is a router object
type RestRoute struct {
	loginController    controller.LoginController
	validateController controller.ValidateController
	loggingMiddleware  gin.HandlerFunc
}

// NewRestRoute init a router object
func NewRestRoute(loginController controller.LoginController, validateController controller.ValidateController, loggingMiddleware gin.HandlerFunc) RestRoute {
	return RestRoute{
		loginController:    loginController,
		validateController: validateController,
		loggingMiddleware:  loggingMiddleware,
	}
}

// Run start a rest server
func (route *RestRoute) Run(serverAddress string) {
	log := logger.Get()
	server := gin.New()
	server.Use(route.loggingMiddleware)
	server.POST("/login", route.loginController.Login)
	server.POST("/validate", route.validateController.Validate)

	err := server.Run(serverAddress)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
		return
	}
}
