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

type PaymentController struct {
	rg             *gin.RouterGroup
	paymentUseCase usecase.PaymentUseCase
	tokenMdw       middleware.AuthTokenMiddleware
}

func (pc *PaymentController) createNewPayment(ctx *gin.Context) {
	newPayment := new(model.Payment)

	if err := ctx.ShouldBindJSON(newPayment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bin JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := pc.paymentUseCase.Insert(newPayment); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: insert or update on table") {
			errorMessage = "sender ID or receiver ID or bank account number does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new payment",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new payment",
	})
}

func (pc *PaymentController) getPayments(ctx *gin.Context) {
	payments, err := pc.paymentUseCase.Payments()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get payments",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get all payments",
		"payments": payments,
	})
}

func (pc *PaymentController) getPaymentById(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse payment ID",
			"error":   err.Error(),
		})
		return
	}

	payment, err := pc.paymentUseCase.PaymentById(paymentId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Payment with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get payment",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get payment by ID",
		"payment": payment,
	})
}

func NewPaymentController(rg *gin.RouterGroup, paymentUseCase usecase.PaymentUseCase, tokenMdw middleware.AuthTokenMiddleware) *PaymentController {
	controller := PaymentController{
		rg:             rg,
		paymentUseCase: paymentUseCase,
		tokenMdw:       tokenMdw,
	}

	protectedGroup := rg.Group("", tokenMdw.RequiredToken())

	protectedGroup.POST("/payments", controller.createNewPayment)
	protectedGroup.GET("/payments", controller.getPayments)
	protectedGroup.GET("/payments/:id", controller.getPaymentById)

	return &controller
}
