package routes

import (
	"go-account-service/application/http/controllers"

	"github.com/gin-gonic/gin"
)

// InscribeAccountRoutes sets the handlers for the accounts routes.
func InscribeAccountRoutes(server *gin.Engine) {
	server.GET("/accounts/:accountId", controllers.GetAccountInfo)
}
