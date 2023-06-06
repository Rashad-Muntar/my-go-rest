package main

import (
	"github.com/Rashad-Muntar/my-go-rest/config"
	"github.com/Rashad-Muntar/my-go-rest/initializers"
	"github.com/Rashad-Muntar/my-go-rest/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVaraiables()
	config.Connect() 
}
func main() {

	r := gin.New()
	routes.UserRoute(r)
	r.Run(":8080")
}
