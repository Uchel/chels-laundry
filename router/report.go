package router

import (
	"database/sql"

	"github.com/Uchel/chels-laundry/controllers"
	"github.com/Uchel/chels-laundry/repositories"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

func ReportRouter(router *gin.Engine, db *sql.DB) {
	reportRepo := repositories.NewReportRepo(db)
	reportUc := usecases.NewReportUsecase(reportRepo)
	reportCtrl := controllers.NewReportController(reportUc)

	reportRouter := router.Group("laundry/employees/report")
	reportRouter.GET("/all-report-detail", reportCtrl.FindAllReportDetail)
	reportRouter.GET("/passcode-report-detail", reportCtrl.FindReportDetailByPassCode)
	reportRouter.GET("/date-report-detail", reportCtrl.FindReportDetailByDate)
	reportRouter.GET("/passcode-report-order", reportCtrl.FindReportOrderByPassscode)
}
