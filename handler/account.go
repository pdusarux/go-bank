package handler

import (
	"encoding/json"
	"go-bank/errs"
	"go-bank/logs"
	"go-bank/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(c *gin.Context) {
	if c.Request.Header.Get("content-type") != "application/json" {
		handleError(c, errs.NewValidationError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		handleError(c, errs.NewValidationError("request body incorrect format"))
		return
	}

	customerID, err := strconv.Atoi(c.Param("customer_id"))
	logs.Info("customerID", zap.Int("customerID", customerID))
	if err != nil {
		handleError(c, err)
		return
	}

	account, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, account)

}

func (h accountHandler) GetAccounts(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		handleError(c, err)
		return
	}

	accounts, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, accounts)
}
