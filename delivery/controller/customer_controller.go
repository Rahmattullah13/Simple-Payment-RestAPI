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

type CustomerController struct {
	rg              *gin.RouterGroup
	customerUseCase usecase.CustomerUseCase
	tokenMdw        middleware.AuthTokenMiddleware
}

// @Summary Get all customers
// @Tags customer
// @success 200 {object} model.CustomerResponse{data=[]model.Customer}
// @Router /api/customers [get]
func (cc *CustomerController) getCustomers(ctx *gin.Context) {

	customers, err := cc.customerUseCase.Customers()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get customers",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all customers",
		"data":    customers,
	})
}

// @Summary Create new customer
// @Tags customer
// @Param string body model.CustomerRequest true "Customer"
// @Success 200 {object} model.CustomerResponse
// @Failure 400
// @Router /api/customers [post]
func (cc *CustomerController) createNewCustomer(ctx *gin.Context) {
	newCustomer := new(model.Customer)

	if err := ctx.ShouldBindJSON(newCustomer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.customerUseCase.Insert(newCustomer); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "User with that ID already a customer"
		}

		if strings.Contains(err.Error(), "pq: insert or update on table") {
			errorMessage = "User with that ID does not exist, Please sign up first"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new customer",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new customer",
	})
}

// @Summary Get customer by ID
// @Tags customer
// @Param id path int true "Customer ID"
// @Success 200 {object} model.Customer
// @Router /api/customers/{id} [get]
func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse customer ID",
			"error":   err.Error(),
		})
		return
	}

	customer, err := cc.customerUseCase.CustomerById(customerId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Customer with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get customer",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get customer by ID",
		"customer": customer,
	})
}

// @Summary Topup customer by ID
// @Tags customer
// @Param object body model.TopUpRequest{balance=int} true "TopUp"
// @Success 200 {object} model.Customer
// @Router /api/customers/id [put]
func (cc *CustomerController) topUpCustomerById(ctx *gin.Context) {
	customer := new(model.Customer)

	if err := ctx.ShouldBindJSON(customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
		})
		return
	}

	if err := cc.customerUseCase.TopUpCustomerBalance(customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to topup customer",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Topup success",
	})
}

// @Summary Delete customer by ID
// @Tags customer
// @Param id path int true "Customer ID"
// @Success 200
// @Router /api/customers/{id} [delete]
func (cc *CustomerController) deleteCustomerById(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse customer ID",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.customerUseCase.Delete(customerId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete customer",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete customer success",
	})
}

func NewCustomerController(rg *gin.RouterGroup, customerUseCase usecase.CustomerUseCase, tokenMdw middleware.AuthTokenMiddleware) *CustomerController {
	controller := CustomerController{
		rg:              rg,
		customerUseCase: customerUseCase,
		tokenMdw:        tokenMdw,
	}

	protectedGroup := rg.Group("", tokenMdw.RequiredToken())

	protectedGroup.GET("/customers", controller.getCustomers)
	protectedGroup.POST("/customers", controller.createNewCustomer)
	protectedGroup.GET("/customers/:id", controller.getCustomerById)
	protectedGroup.PUT("/customers/:id", controller.topUpCustomerById)
	protectedGroup.DELETE("/customers/:id", controller.deleteCustomerById)

	return &controller
}
