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

type BankController struct {
	rg          *gin.RouterGroup
	bankUseCase usecase.BankUseCase
	tokenMdw    middleware.AuthTokenMiddleware
}

// @Summary Get all banks
// @Tags Bank endpoints
// @success 200 {object} model.BankResponse{data=[]model.Bank}
// @Router /api/banks [get]
func (bc *BankController) getBanks(ctx *gin.Context) {

	banks, err := bc.bankUseCase.Banks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get banks",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all banks",
		"banks":   banks,
	})
}

// @Summary Create new bank
// @Tags Bank endpoints
// @Param string body model.BankRequest true "Bank"
// @Success 200 {object} model.BankResponse
// @Failure 400
// @Router /api/banks [post]
func (bc *BankController) createNewBank(ctx *gin.Context) {
	newBank := new(model.Bank)

	if err := ctx.ShouldBindJSON(newBank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := bc.bankUseCase.Insert(newBank); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "Bank with that account number already exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new bank",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new bank",
	})
}

// @Summary Get bank by ID
// @Tags Bank endpoints
// @Param id path int true "Bank ID"
// @Success 200 {object} model.Bank
// @Router /api/banks/{id} [get]
func (bc *BankController) getBankById(ctx *gin.Context) {
	bankId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse bank ID",
			"error":   err.Error(),
		})
		return
	}

	bank, err := bc.bankUseCase.BankById(bankId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Bank with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get bank",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get bank by ID",
		"banks":   bank,
	})
}

// @Summary Delete bank by ID
// @Tags Bank endpoints
// @Param id path int true "Bank ID"
// @Success 200
// @Router /api/banks/{id} [delete]
func (bc *BankController) deleteBankById(ctx *gin.Context) {
	bankId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse bank ID",
			"error":   err.Error(),
		})
		return
	}

	if err := bc.bankUseCase.Delete(bankId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete bank",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete bank success",
	})
}

func NewBankController(rg *gin.RouterGroup, bankUseCase usecase.BankUseCase, tokenMdw middleware.AuthTokenMiddleware) *BankController {
	controller := BankController{
		rg:          rg,
		bankUseCase: bankUseCase,
		tokenMdw:    tokenMdw,
	}

	protectedGroup := rg.Group("", tokenMdw.RequiredToken())

	protectedGroup.GET("/banks", controller.getBanks)
	protectedGroup.POST("/banks", controller.createNewBank)
	protectedGroup.GET("/banks/:id", controller.getBankById)
	protectedGroup.DELETE("/banks/:id", controller.deleteBankById)

	return &controller
}
