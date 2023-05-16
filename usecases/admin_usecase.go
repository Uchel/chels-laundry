package usecases

import (
	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/repositories"
)

type AdminLaundryUsecase interface {
	FindAll() any
	FindByEmail(email string) any
	FindPhotoByEmail(email string) any
	Register(newAdmin *models.Admin) string
	Edit(admin *models.Admin, email string) string
	EditPhoto(photo string, email string) string
	Unregister(email string) string
}

type adminLaundryUsecase struct {
	adminLaundryRepo repositories.AdminLaundryRepo
}

// ============================================= Find All===============================================
func (u *adminLaundryUsecase) FindAll() any {
	return u.adminLaundryRepo.GetAll()
}

// ============================================= Find By Email===============================================
func (u *adminLaundryUsecase) FindByEmail(email string) any {
	return u.adminLaundryRepo.GetByEmail(email)
}

// ============================================= Find Photo By Email===============================================
func (u *adminLaundryUsecase) FindPhotoByEmail(email string) any {
	return u.adminLaundryRepo.GetPhotoByEmail(email)
}

// ============================================= Register ===============================================
func (u *adminLaundryUsecase) Register(newAdmin *models.Admin) string {
	return u.adminLaundryRepo.Create(newAdmin)
}

// ============================================= Edit ===============================================
func (u *adminLaundryUsecase) Edit(admin *models.Admin, email string) string {
	return u.adminLaundryRepo.Update(admin, email)
}

// ============================================= Edit Photo ===============================================
func (u *adminLaundryUsecase) EditPhoto(photo string, email string) string {
	return u.adminLaundryRepo.UpdatePhoto(photo, email)
}

// ============================================= UnRegister ===============================================
func (u *adminLaundryUsecase) Unregister(email string) string {
	return u.adminLaundryRepo.Delete(email)
}

func NewAdminLaundryUC(adminLaundryRepo repositories.AdminLaundryRepo) AdminLaundryUsecase {
	return &adminLaundryUsecase{
		adminLaundryRepo: adminLaundryRepo,
	}
}
