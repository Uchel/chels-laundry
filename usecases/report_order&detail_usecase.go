package usecases

import (
	"fmt"

	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/repositories"
)

type ReportUsecase interface {
	FindAllReportDetail() any
	FindReportDetailByPassCode(passcode *models.DetailReportReq) any
	FindReportDetailByDate(dateout *models.DetailReportReq) any
	FindReportOrderByPassscode(passcode *models.OrderReport) any
}

type reportUsecase struct {
	reportRepo repositories.ReportRepo
}

func (u *reportUsecase) FindAllReportDetail() any {

	return u.reportRepo.GetAllReportDetail()
}
func (u *reportUsecase) FindReportDetailByPassCode(passcode *models.DetailReportReq) any {
	return u.reportRepo.GetReportDetailByPassCode(passcode)
}
func (u *reportUsecase) FindReportDetailByDate(dateout *models.DetailReportReq) any {
	return u.reportRepo.GetReportDetailByDate(dateout)
}
func (u *reportUsecase) FindReportOrderByPassscode(passcode *models.OrderReport) any {
	fmt.Println(passcode)
	return u.reportRepo.GetReportOrderByPassscode(passcode)
}

func NewReportUsecase(reportRepo repositories.ReportRepo) ReportUsecase {
	return &reportUsecase{
		reportRepo: reportRepo,
	}
}
