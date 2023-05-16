package usecases

import (
	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/repositories"
)

type TrxOrderListUsecases interface {
	TrxEnrollOrderList(order *models.OrderList) (string, any)
	TrxOrderConfirmation(report *models.DetailReportReq) string
	OrderCancel() string
	FinishingTrxOrder(report *models.DetailReportReq) string
}

type trxOrderListUsecases struct {
	trxOrderListRepo repositories.TrxOrderListRepo
}

// =======================================================================================================
func (u trxOrderListUsecases) TrxEnrollOrderList(order *models.OrderList) (string, any) {
	return u.trxOrderListRepo.TrxOrderList(order)
}
func (u trxOrderListUsecases) TrxOrderConfirmation(report *models.DetailReportReq) string {
	return u.trxOrderListRepo.OrderConfirmation(report)
}
func (u trxOrderListUsecases) OrderCancel() string {
	return u.trxOrderListRepo.OrderCancel()
}

func (u trxOrderListUsecases) FinishingTrxOrder(report *models.DetailReportReq) string {
	return u.trxOrderListRepo.FinishingOrder(report)
}

// =======================================================================================================
func NewTrxOrderListUC(trxOrderListRepo repositories.TrxOrderListRepo) TrxOrderListUsecases {
	return &trxOrderListUsecases{
		trxOrderListRepo: trxOrderListRepo,
	}
}
