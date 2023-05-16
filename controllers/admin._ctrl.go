package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/usecases"
	"github.com/gin-gonic/gin"
)

type AdminLaundryCtrl struct {
	usecases usecases.AdminLaundryUsecase
}

// ==================================================== Register =========================================
func (c *AdminLaundryCtrl) Register(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// email := claims["email"].(string)

	var newAdmin models.Admin

	if err := ctx.ShouldBind(&newAdmin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Photo"})
		return
	}

	//ambil filename untuk isi string foto
	filename := file.Filename
	newAdmin.Photo = filename

	if errFile := ctx.SaveUploadedFile(file, fmt.Sprintf("./delivery/images/%s", filename)); errFile != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := c.usecases.Register(&newAdmin)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
		// "login_with": email,
	})

}

// ==================================================== Find All  ====================================================
func (c *AdminLaundryCtrl) FindAll(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// role := claims["role"].(string)
	// if role != "ic" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "you have no access to this role",
	// 	})
	// 	return
	// }

	res := c.usecases.FindAll()

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
		// "login_with": email,
	})

}

// ==================================================== Find AdminPhoto by email =========================================
func (c *AdminLaundryCtrl) FindPhotoByEmail(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// role := claims["role"].(string)
	// if role != "ic" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "you have no access to this role",
	// 	})
	// 	return
	// }

	email := ctx.Param("email")

	photo := c.usecases.FindPhotoByEmail(email)
	filepath := fmt.Sprintf("./delivery/images/%s", photo)

	file, err := os.Open(filepath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", fileBytes)
}

// ====================================================Find Data By Email================================================
func (c *AdminLaundryCtrl) FindByEmail(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// role := claims["role"].(string)
	// if role != "ic" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "you have no access to this role",
	// 	})
	// 	return
	// }

	email := ctx.Param("email")

	res := c.usecases.FindByEmail(email)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
		// "login_with": email,
	})

}

// ==================================================== Edit Photo ================================================

func (c *AdminLaundryCtrl) EditPhoto(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// role := claims["role"].(string)
	// if role != "ic" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "you have no access to this role",
	// 	})
	// 	return
	// }

	email := ctx.Param("email")
	var admin models.Admin
	//=============================== Delete old Photo ===============================
	filename1 := c.usecases.FindPhotoByEmail(email)

	err := os.Remove(fmt.Sprintf("./delivery/images/%s", filename1))
	if err != nil {
		log.Println(err)
	}
	//============================ update new photo ==================================
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Photo"})
		return
	}

	//ambil filename untuk isi string foto
	filename := file.Filename
	admin.Photo = filename

	if errFile := ctx.SaveUploadedFile(file, fmt.Sprintf("./delivery/images/%s", filename)); errFile != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	res := c.usecases.EditPhoto(admin.Photo, email)
	if err := ctx.ShouldBind(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
		// "login_with": email,
	})

}

// ==================================================== Edit Data ================================================

func (c *AdminLaundryCtrl) EditPassword(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// role := claims["role"].(string)
	// if role != "ic" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "you have no access to this role",
	// 	})
	// 	return
	// }

	email := ctx.Param("email")
	var admin *models.Admin
	//=============================== Delete old Photo ===============================

	//ambil filename untuk isi string foto

	if err := ctx.ShouldBind(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	res := c.usecases.Edit(admin, email)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
		// "login_with": email,
	})

}
func (c *AdminLaundryCtrl) DeleteByEmail(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// email := claims["email"].(string)
	// role := claims["role"].(string)
	// if role != "wh" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "you have no access to this role",
	// 	})
	// 	return
	// }

	email := ctx.Param("email")

	filename := c.usecases.FindByEmail(email)

	res := c.usecases.Unregister(email)
	if res == "admin not found" {
		ctx.JSON(http.StatusBadRequest, "invalid input ID")
		return
	}

	if filename != "" {
		err := os.Remove(fmt.Sprintf("./images/%s", filename))
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "deleted success"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
		// "login with": email,
	})

}

func NewAdminLaundryCtrl(c usecases.AdminLaundryUsecase) *AdminLaundryCtrl {
	controller := AdminLaundryCtrl{
		usecases: c,
	}
	return &controller
}
