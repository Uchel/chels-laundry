package router

import (
	"database/sql"

	"github.com/Uchel/chels-laundry/controllers"
	"github.com/Uchel/chels-laundry/repositories"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

func RouterTrxOrderList(router *gin.Engine, db *sql.DB) {

	// jwtAdminIcRepo := jwt_repository.NewIcTeamLoginRepo(db)
	// jwtIcUsecase := jwt_usecase.NewIcTeamUsecase(jwtAdminIcRepo)
	// jwtIcCtrl := jwt_controller.NewIcTeamLoginController(jwtIcUsecase, 60)

	trxOrderListRepo := repositories.NewOrderListRepo(db)
	trxOrderListUsecase := usecases.NewTrxOrderListUC(trxOrderListRepo)
	trxOrderListCtrl := controllers.NewTrxOrderListCtrl(trxOrderListUsecase)

	//============================= Warehouse/SuperAdmin akses(register dan delete ic_team/admin_id) ===============================

	TrxOrderListRouter := router.Group("/laundry/employees/trx")
	TrxOrderListRouter.POST("/order-list", trxOrderListCtrl.Register)
	TrxOrderListRouter.POST("/confirmation", trxOrderListCtrl.OrderConfirmation)
	TrxOrderListRouter.DELETE("/cancel", trxOrderListCtrl.OrderCancel)
	TrxOrderListRouter.PUT("/finishing", trxOrderListCtrl.FinishingTrxOrder)

	//=========================================================================================================================

}
