package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gihub.com/eugenenosenko/bookstore-users-api/domain/users"
	"gihub.com/eugenenosenko/bookstore-users-api/services"
	"gihub.com/eugenenosenko/bookstore-users-api/utils/errors"
)

func GetUser(ctx *gin.Context) {
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.GetUser(userID)

	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func CreateUsers(ctx *gin.Context) {
	var user users.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func SearchUsers(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "this endpoint is not implemented")
}
