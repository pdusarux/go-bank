package handler

import (
	"go-bank/errs"
	"go-bank/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context, err error) {
	logs.Error(err)
	switch e := err.(type) {
	case errs.AppError:
		ctx.JSON(e.Code, gin.H{"error": e.Message})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
