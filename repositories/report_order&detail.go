package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Uchel/chels-laundry/models"
)

type ReportRepo interface {
	GetAllReportDetail() any
	GetReportDetailByPassCode(passcode *models.DetailReportReq) any
	GetReportDetailByDate(dateout *models.DetailReportReq) any
	GetReportOrderByPassscode(passcode *models.OrderReport) any
}

type reportRepo struct {
	db *sql.DB
}

func (r *reportRepo) GetAllReportDetail() any {
	var reportsDetail []models.DetailReportReq

	query := "SELECT pass_code, paid_status,date_in,date_out,taken_by,total_price from report_detail order by date_out asc"

	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var report models.DetailReportReq

		if err := rows.Scan(&report.Pascode, &report.PaidStatus, &report.DateIn, &report.DateOut, &report.TakenBy, &report.TotalPrice); err != nil {
			log.Println(err)
		}

		reportsDetail = append(reportsDetail, report)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return "internal server error"
	}

	if len(reportsDetail) == 0 {
		return "no data"
	}

	return reportsDetail
}

func (r *reportRepo) GetReportDetailByPassCode(passcode *models.DetailReportReq) any {
	query := "SELECT pass_code, paid_status,date_in,date_out,taken_by,total_price from report_detail where pass_code = $1"
	row := r.db.QueryRow(query, &passcode.Pascode)

	if err := row.Scan(&passcode.Pascode, &passcode.PaidStatus, &passcode.DateIn, &passcode.DateOut, &passcode.TakenBy, &passcode.TotalPrice); err != nil {
		log.Println(err)
	}
	if passcode.Pascode == "" {
		return "passcode not found"
	}
	return passcode
}

func (r *reportRepo) GetReportDetailByDate(dateout *models.DetailReportReq) any {
	var reportsDetail []models.DetailReportReq
	query := "SELECT pass_code, paid_status,date_in,date_out,taken_by,total_price from report_detail where date_out = $1"
	rows, err := r.db.Query(query, dateout.DateOut)
	fmt.Println(dateout.DateOut, "oyy")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var report models.DetailReportReq

		if err := rows.Scan(&report.Pascode, &report.PaidStatus, &report.DateIn, &report.DateOut, &report.TakenBy, &report.TotalPrice); err != nil {
			log.Println(err)
		}

		reportsDetail = append(reportsDetail, report)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsDetail) == 0 {
		return "no data or order not found"
	}

	return reportsDetail
}

func (r *reportRepo) GetReportOrderByPassscode(passcode *models.OrderReport) any {
	var reportsOrder []models.OrderReport
	query := "select order_list_id, name,pass_code,service,quantity,unit,total,created_at from report_order where pass_code = $1"

	rows, err := r.db.Query(query, passcode.PassCode)
	fmt.Println(passcode.PassCode, "oyy")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var report models.OrderReport

		if err := rows.Scan(&report.OrderListId, &report.Name, &report.PassCode, &report.Service, &report.Quantity, &report.Unit, &report.Total, &report.Created_At); err != nil {
			log.Println(err)
		}

		reportsOrder = append(reportsOrder, report)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsOrder) == 0 {
		return "no data or order not found"
	}

	return reportsOrder
}

func NewReportRepo(db *sql.DB) ReportRepo {
	repo := new(reportRepo)
	repo.db = db
	return repo
}
