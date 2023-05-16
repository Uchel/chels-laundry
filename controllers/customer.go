package controllers

import (
	"net/http"

	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

type CustomerCtrl struct {
	usecases usecases.CustomerUsecases
}

func (c *CustomerCtrl) Register(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// email := claims["email"].(string)

	var newCustomer models.Customer

	if err := ctx.ShouldBind(&newCustomer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	res := c.usecases.GetPasscode(&newCustomer)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
		// "login_with": email,
	})

}

func NewCustomerCtrl(c usecases.CustomerUsecases) *CustomerCtrl {
	controller := CustomerCtrl{
		usecases: c,
	}
	return &controller
}
