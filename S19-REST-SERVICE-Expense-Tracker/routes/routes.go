package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shanker.com/expense-tracker/models"
)

func RegisterRoutes(server *gin.Engine){

	server.GET("/expense",getAllExpenses)

	server.POST("/expense",createExpense)

	server.PUT("/expense/:id",editExpense)

	server.DELETE("/expense/:id",deleteExpense)
}


func getAllExpenses(context *gin.Context) {
	expenses , err := models.GetAllExpenses()
	
	if err!=nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt Fetch!!"})
		return
	}

	context.JSON(http.StatusOK,expenses)
}


func createExpense(context *gin.Context){

	var expense models.Expense
	err := context.ShouldBindJSON(&expense);

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"couldnt parse request"})
		return
	}

	err = expense.Save()
	
	
	if(err!= nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt Create expense, Try again later!!"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message":"Expense  Created!","event":expense})
}

func editExpense(context *gin.Context){

	var expense models.Expense
	err := context.ShouldBindJSON(&expense);
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"couldnt parse request"})
		return
	}

	eId,err := strconv.ParseInt(context.Param("id"),10,64) 

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"couldnt parse request (PARAM)"})
		return
	}

	err = expense.Modify(eId)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Some error occured!!"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Record Updated!!!!"})

}


func deleteExpense(context *gin.Context){
	eId , err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldnt parse expense ID"})
		return
	}
	
	err = models.Delete(eId)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Couldnt delete the record!"})
		return
	}

	context.JSON(http.StatusBadRequest, gin.H{"message":"Deleted the record!!"})

}