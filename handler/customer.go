package handler

import (
	"go-bank/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) GetCustomers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := h.custSrv.GetCustomers()
		if err != nil {
			handleError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, customers)
	}
}

func (h customerHandler) GetCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerID, err := strconv.Atoi(ctx.Param("customer_id"))
		if err != nil {
			handleError(ctx, err)
			return
		}

		customer, err := h.custSrv.GetCustomer(customerID)
		if err != nil {
			handleError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, customer)
	}
}
