package router

import (
	"database/sql"

	"github.com/Uchel/chels-laundry/controllers"
	"github.com/Uchel/chels-laundry/repositories"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

func RouterCustomer(router *gin.Engine, db *sql.DB) {

	// jwtAdminIcRepo := jwt_repository.NewIcTeamLoginRepo(db)
	// jwtIcUsecase := jwt_usecase.NewIcTeamUsecase(jwtAdminIcRepo)
	// jwtIcCtrl := jwt_controller.NewIcTeamLoginController(jwtIcUsecase, 60)

	customerRepo := repositories.NewCustomerRepo(db)
	customerUsecase := usecases.NewCustomerUC(customerRepo)
	customerCtrl := controllers.NewCustomerCtrl(customerUsecase)

	//============================= Warehouse/SuperAdmin akses(register dan delete ic_team/admin_id) ===============================

	TrxOrderListRouter := router.Group("/laundry/employees/register")
	TrxOrderListRouter.POST("/customer", customerCtrl.Register)

	//=========================================================================================================================

}
