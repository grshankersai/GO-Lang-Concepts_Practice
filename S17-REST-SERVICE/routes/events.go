package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shanker.com/rest-service/models"
)


func getEvents(context *gin.Context){
	events , err:= models.GetAllEvents()
	if err!=nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt fetch, Try again later"})
		return
	}
	// context.JSON(http.StatusOK, gin.H{"message":"Hello! Server is running"})
	context.JSON(http.StatusOK, events)
}


func getEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("eventId"),10,64) 

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt parse event ID"})
		return
	}

	event , err := models.GetEventById(eventId)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt fetch event"})
		return
	}

	context.JSON(http.StatusOK,event)

}

func createEvent(context *gin.Context){

	// token := context.Request.Header.Get("Authorization")

	// if token == "" {
	// 	context.JSON(http.StatusUnauthorized,gin.H{"message":"Not authorized"})
	// 	return
	// }

	// userId,err := utils.VerifyToken(token)

	// if err != nil{
	// 	context.JSON(http.StatusUnauthorized,gin.H{"message":"Not Authorized."})
	// 	return
	// }

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"couldnt parse request"})
		return
	}

	// event.ID = 1
	
	event.UserId = context.GetInt64("userId")

	err = event.Save()	
	if(err!= nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt Create event, Try again later"})
	}

	context.JSON(http.StatusCreated,gin.H{"message":"Event Created!","event":event})
	
}


func updateEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("eventId"),10,64)
	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt parse event ID"})
		return
	}

	uId := context.GetInt64("userId")
	event , err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt fetch event"})
		return
	}

	if(event.UserId != uId){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not authorized to update event!!"})
		return 
	}



	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt parse event data"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if(err!= nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt Create event, Try again later"})
	}

	context.JSON(http.StatusOK,gin.H{"message":"Event Updated Sucessfully!"})



}


func deleteEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("eventId"),10,64)
	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt parse event ID"})
		return
	}


	uId := context.GetInt64("userId")
	event , err := models.GetEventById(eventId)
	if(event.UserId != uId){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not authorized to delete event!!"})
		return 
	}


	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt delete event"})
		return
	}

	err = event.Delete()

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Event Deleted Sucessfully"})




}