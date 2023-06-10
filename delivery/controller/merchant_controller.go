package controller

import (
	"net/http"
	"simple-payment/delivery/middleware"
	"simple-payment/model"
	"simple-payment/usecase"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	rg              *gin.RouterGroup
	merchantUseCase usecase.MerchantUseCase
	tokenMdw        middleware.AuthTokenMiddleware
}

func (mc *MerchantController) createNewMerchant(ctx *gin.Context) {
	newMerchant := new(model.Merchant)

	if err := ctx.ShouldBindJSON(newMerchant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := mc.merchantUseCase.Insert(newMerchant); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "User with that ID already have a merchant"
		}

		if strings.Contains(err.Error(), "pq: insert or update on table") {
			errorMessage = "User with that ID does not exist, Please sign up first"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new merchant",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new merchant",
	})
}

func (mc *MerchantController) getMerchants(ctx *gin.Context) {
	merchants, err := mc.merchantUseCase.Merchants()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get merchants",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get all merchants",
		"merchants": merchants,
	})
}

func (mc *MerchantController) getMerchantById(ctx *gin.Context) {
	merchantId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse merchant ID",
			"error":   err.Error(),
		})
		return
	}

	merchant, err := mc.merchantUseCase.MerchantById(merchantId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Merchant with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get merchant",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get merchant by ID",
		"customer": merchant,
	})
}

func (cc *MerchantController) deleteMerchantById(ctx *gin.Context) {
	merchantId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse merchant ID",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.merchantUseCase.Delete(merchantId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete merchant",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete merchant success",
	})
}

func NewMerchantController(rg *gin.RouterGroup, merchantUseCase usecase.MerchantUseCase, tokenMdw middleware.AuthTokenMiddleware) *MerchantController {
	controller := MerchantController{
		rg:              rg,
		merchantUseCase: merchantUseCase,
		tokenMdw:        tokenMdw,
	}

	protectedGroup := rg.Group("", tokenMdw.RequiredToken())

	protectedGroup.POST("/merchants", controller.createNewMerchant)
	protectedGroup.GET("/merchants", controller.getMerchants)
	protectedGroup.GET("/merchants/:id", controller.getMerchantById)
	protectedGroup.DELETE("/merchants/:id", controller.deleteMerchantById)

	return &controller
}
