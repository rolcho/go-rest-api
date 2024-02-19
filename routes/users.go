package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/models"
)

func getUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func getUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch user id"})
		return
	}
	user, err := models.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch user id"})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func createUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	userExist, err := models.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading the database"})
		return
	}

	if userExist != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "User already exsist"})
		return
	}

	if err := user.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error writing the database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "user": &user})
}

func updateUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch user id"})
		return
	}
	_, err = models.GetUserById(userId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch user id"})
		return
	}

	var updatedUser models.User

	if err = ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	updatedUser.Id = userId

	userExist, err := models.GetUserByEmail(updatedUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading the database"})
		return
	}

	if userExist != nil && userExist.Id != userId {
		ctx.JSON(http.StatusConflict, gin.H{"message": "User already exsist"})
		return
	}

	if err = updatedUser.Update(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}

func deleteUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch user id"})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch user id"})
		return
	}

	if err := user.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
