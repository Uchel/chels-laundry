package delivery

import (
	"github.com/Uchel/chels-laundry/config"
	"github.com/Uchel/chels-laundry/router"
	"github.com/Uchel/chels-laundry/utils"

	"github.com/gin-gonic/gin"
)

func Exec() {
	r := gin.Default()
	db := config.ConnectDB()

	router.RouterAdminLaundry(r, db)
	router.RouterTrxOrderList(r, db)
	router.RouterCustomer(r, db)
	router.ReportRouter(r, db)
	r.Run(":" + utils.DotEnv("PORT"))
}
