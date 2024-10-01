package main

import (
	"github.com/gin-gonic/gin"
	"shanker.com/expense-tracker/db"
	"shanker.com/expense-tracker/routes"
)

func main(){
	server := gin.Default()

	db.InitDB()

	routes.RegisterRoutes(server)

	server.Run(":8081")
}