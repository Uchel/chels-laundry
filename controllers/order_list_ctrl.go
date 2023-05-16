package controllers

import (
	"net/http"

	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

type TrxOrderListCtrl struct {
	usecases usecases.TrxOrderListUsecases
}

func (c *TrxOrderListCtrl) Register(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// email := claims["email"].(string)

	var newOrder models.OrderList

	if err := ctx.ShouldBind(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	res, data := c.usecases.TrxEnrollOrderList(&newOrder)

	ctx.JSON(http.StatusCreated, gin.H{
		"respon": res,
		"data":   data,
		// "login_with": email,
	})

}

// =================================================================confirmation===========================
func (c *TrxOrderListCtrl) OrderConfirmation(ctx *gin.Context) {
	var report models.DetailReportReq

	if err := ctx.ShouldBind(&report); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	res := c.usecases.TrxOrderConfirmation(&report)
	ctx.JSON(http.StatusCreated, gin.H{
		"respon": res,

		// "login_with": email,
	})
}
func (c *TrxOrderListCtrl) OrderCancel(ctx *gin.Context) {

	res := c.usecases.OrderCancel()
	ctx.JSON(http.StatusCreated, gin.H{
		"respon": res,

		// "login_with": email,
	})
}

func (c *TrxOrderListCtrl) FinishingTrxOrder(ctx *gin.Context) {
	var report models.DetailReportReq

	if err := ctx.ShouldBind(&report); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	res := c.usecases.FinishingTrxOrder(&report)
	ctx.JSON(http.StatusCreated, gin.H{
		"respon": res,

		// "login_with": email,
	})
}

func NewTrxOrderListCtrl(c usecases.TrxOrderListUsecases) *TrxOrderListCtrl {
	controller := TrxOrderListCtrl{
		usecases: c,
	}
	return &controller
}
