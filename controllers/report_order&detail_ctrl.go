package controllers

import (
	"fmt"
	"net/http"

	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUsecase usecases.ReportUsecase
}

func (c ReportController) FindAllReportDetail(ctx *gin.Context) {
	res := c.reportUsecase.FindAllReportDetail()
	if res == "internal server error" {
		ctx.JSON(500, res)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}
func (c ReportController) FindReportDetailByPassCode(ctx *gin.Context) {
	var passcode models.DetailReportReq
	if err := ctx.ShouldBindJSON(&passcode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}
	fmt.Println(passcode)
	res := c.reportUsecase.FindReportDetailByPassCode(&passcode)
	ctx.JSON(200, gin.H{
		"data": res,
	})
}

func (c ReportController) FindReportDetailByDate(ctx *gin.Context) {

	var dateout models.DetailReportReq
	if err := ctx.ShouldBindJSON(&dateout); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}
	fmt.Println(dateout)
	res := c.reportUsecase.FindReportDetailByDate(&dateout)
	ctx.JSON(200, gin.H{
		"data": res,
	})
}

func (c ReportController) FindReportOrderByPassscode(ctx *gin.Context) {
	var passcode models.OrderReport
	if err := ctx.ShouldBindJSON(&passcode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	res := c.reportUsecase.FindReportOrderByPassscode(&passcode)
	ctx.JSON(200, gin.H{
		"data": res,
	})
}

func NewReportController(ctrl usecases.ReportUsecase) *ReportController {
	controller := ReportController{
		reportUsecase: ctrl,
	}

	return &controller
}
