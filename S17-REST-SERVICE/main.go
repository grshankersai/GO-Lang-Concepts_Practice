package main

import (
	"github.com/gin-gonic/gin"
	"shanker.com/rest-service/db"
	"shanker.com/rest-service/routes"
)



func main(){
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)


	server.Run(":8080")

}

	