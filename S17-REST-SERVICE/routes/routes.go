package routes

import (
	"github.com/gin-gonic/gin"
	"shanker.com/rest-service/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events",getEvents)
	server.GET("/events/:eventId",getEvent)


	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events" , createEvent)
	authenticated.PUT("/events/:eventId",updateEvent)
	authenticated.DELETE("/events/:eventId",deleteEvent)

	authenticated.POST("/events/:eventId/register",registerForEvent)
	authenticated.DELETE("/events/:eventId/register",cancelRegistration)

	// server.POST("/events" , middlewares.Authenticate , createEvent)

	// server.PUT("/events/:eventId",updateEvent)

	// server.DELETE("/events/:eventId",deleteEvent)


	server.POST("/signup",signup)

	server.POST("/login",login)

	// list all registrations
}