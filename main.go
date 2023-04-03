package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joaopaulo-bertoncini/jwt-auth-gin/src/controller"
	"github.com/joaopaulo-bertoncini/jwt-auth-gin/src/service"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
