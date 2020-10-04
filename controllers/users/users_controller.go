package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gihub.com/eugenenosenko/bookstore-users-api/domain/users"
	"gihub.com/eugenenosenko/bookstore-users-api/services"
	"gihub.com/eugenenosenko/bookstore-users-api/utils/errors"
)

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "this endpoint is not implemented")
}

func CreateUsers(ctx *gin.Context) {
	var user users.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := &errors.RestErr{
			Message: "invalid json",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
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
