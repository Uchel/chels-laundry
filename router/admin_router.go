package router

import (
	"database/sql"

	"github.com/Uchel/chels-laundry/controllers"
	"github.com/Uchel/chels-laundry/repositories"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

func RouterAdminLaundry(router *gin.Engine, db *sql.DB) {

	// jwtAdminIcRepo := jwt_repository.NewIcTeamLoginRepo(db)
	// jwtIcUsecase := jwt_usecase.NewIcTeamUsecase(jwtAdminIcRepo)
	// jwtIcCtrl := jwt_controller.NewIcTeamLoginController(jwtIcUsecase, 60)

	adminLaundryRepo := repositories.NewExampleAdminRepo(db)
	adminLaundryUsecase := usecases.NewAdminLaundryUC(adminLaundryRepo)
	adminLaundryCtrl := controllers.NewAdminLaundryCtrl(adminLaundryUsecase)

	//============================= Warehouse/SuperAdmin akses(register dan delete ic_team/admin_id) ===============================

	adminLaundryRouter := router.Group("/laundry/employees")
	adminLaundryRouter.POST("/register", adminLaundryCtrl.Register)
	adminLaundryRouter.GET("/photo/:email", adminLaundryCtrl.FindPhotoByEmail)
	adminLaundryRouter.GET("/data/:email", adminLaundryCtrl.FindByEmail)
	adminLaundryRouter.GET("/findall", adminLaundryCtrl.FindAll)
	adminLaundryRouter.PUT("/update_photo/:email", adminLaundryCtrl.EditPhoto)       //change foto
	adminLaundryRouter.PUT("/update_password/:email", adminLaundryCtrl.EditPassword) //change password
	adminLaundryRouter.DELETE("/delete/:email", adminLaundryCtrl.DeleteByEmail)
	//=========================================================================================================================

}
