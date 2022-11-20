package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gzim07/jwt_auth/controllers"
	"github.com/gzim07/jwt_auth/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()

}
func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.Run(":" + port)
}
