package controller

import (
	"net/http"
	"simple-payment/model"
	"simple-payment/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	rg          *gin.RouterGroup
	userUseCase usecase.UserUseCase
}

func (uc *UserController) createNewUser(ctx *gin.Context) {
	newUser := new(model.User)

	if err := ctx.ShouldBindJSON(newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := uc.userUseCase.Insert(newUser); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "User with that email already exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new user",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new user",
	})
}

func (uc *UserController) loginUser(ctx *gin.Context) {
	var userCredential model.UserCredential
	err := ctx.ShouldBindJSON(&userCredential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := uc.userUseCase.Login(userCredential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("token", token, 60*60, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "You are logged in",
		"account": token,
	})
}

func (uc *UserController) logoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "You are logged out",
	})
}

func NewUserController(rg *gin.RouterGroup, userUseCase usecase.UserUseCase) *UserController {
	controller := UserController{
		rg:          rg,
		userUseCase: userUseCase,
	}

	controller.rg.POST("/auth/users", controller.createNewUser)
	controller.rg.POST("/session", controller.loginUser)
	controller.rg.POST("/session/logout", controller.logoutUser)

	return &controller
}
