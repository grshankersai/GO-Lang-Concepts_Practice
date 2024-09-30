package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shanker.com/rest-service/models"
	"shanker.com/rest-service/utils"
)

func signup(context *gin.Context){

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"couldnt parse request"})
		return
	}

	err = user.Save()

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldnt Create event, Try again later"})
	}

	context.JSON(http.StatusCreated,gin.H{"message":"User Created sucessfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"couldnt parse request"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Invalid Credentials"})
		return
	}

	// fmt.Print(user.ID)

	token , err := utils.GenerateToken(user.Email,user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could Not authenticate user."})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Login Success!!!", "token":token})

}

