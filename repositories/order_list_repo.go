package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Uchel/chels-laundry/models"
)

type TrxOrderListRepo interface {
	TrxOrderList(order *models.OrderList) (string, any)
	OrderConfirmation(report *models.DetailReportReq) string
	OrderCancel() string
	FinishingOrder(report *models.DetailReportReq) string
}

type trxOrderListRepo struct {
	db *sql.DB
}

func NewOrderListRepo(db *sql.DB) TrxOrderListRepo {
	repo := new(trxOrderListRepo)
	repo.db = db
	return repo

}

// ===================================================================Order=======================================================
func (r trxOrderListRepo) TrxOrderList(order *models.OrderList) (string, any) {
	id := order.ID
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err.Error(), nil
	}
	// fill query

	r.InsertOrder(order, tx)
	r.UpdateTotalOrderList(tx)
	report := r.GetReport(id, tx)
	r.InsertReportOrder(&report, tx)
	orders := r.GetAllOrder(tx)
	total, _ := r.GetPasscodeTotal(tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err.Error(), "no data"
	} else {
		return fmt.Sprintf("Order Successfully, Total = %d", total), orders
	}
}

// ===================================================================Confirmation ======================================================
func (r trxOrderListRepo) OrderConfirmation(report *models.DetailReportReq) string {

	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	// fill query
	total, passcode := r.GetPasscodeTotal(tx)
	report.Pascode = passcode
	report.TotalPrice = total

	totalPrice := r.InsertReportDetail(report, tx)
	r.TruncateOrder(tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err.Error()
	} else {
		return totalPrice
	}
}

// =======================================================================Cancel ==========================================================
func (r trxOrderListRepo) OrderCancel() string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	// fill query
	_, passcode := r.GetPasscodeTotal(tx)
	r.DeleteReportByPasscode(passcode, tx)
	r.TruncateOrder(tx)
	r.DeleteCustomerByPasscode(passcode, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err.Error()
	} else {
		return "Order Canceled"
	}
}

// ======================================================================= Finishing order ==========================================================
func (r trxOrderListRepo) FinishingOrder(report *models.DetailReportReq) string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	// fill query
	passcode := r.FinishingReportDetail(report, tx)
	r.DeleteCustomerByPasscode(passcode, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err.Error()
	} else {
		return fmt.Sprintf("Order with passcode pascode = %s Finished", passcode)
	}
}

// ==================================== To Finishing=====================================================================================
func (r trxOrderListRepo) FinishingReportDetail(report *models.DetailReportReq, tx *sql.Tx) string {
	query := "update report_detail set date_out = $1 , taken_by = $2, paid_status = $3 where pass_code =$4"
	_, err := tx.Exec(query, time.Now(), &report.TakenBy, true, &report.Pascode)
	Validate(err, "update", tx)
	return report.Pascode
}

// ===============================================================To Cancel =========================================================
func (r trxOrderListRepo) DeleteReportByPasscode(passcode string, tx *sql.Tx) {

	query := "DELETE FROM report_order  WHERE pass_code = $1"
	_, err := tx.Exec(query, passcode)
	Validate(err, "Delete", tx)
}
func (r trxOrderListRepo) DeleteCustomerByPasscode(passcode string, tx *sql.Tx) {

	query := "DELETE FROM customer  WHERE pass_code = $1"
	_, err := tx.Exec(query, passcode)
	Validate(err, "Delete", tx)
}

// ==============================================================To Confirmation and interim report ===========================================
// =============================================================Get passcode and sum total price order list============================

func (r trxOrderListRepo) GetPasscodeTotal(tx *sql.Tx) (int, string) {
	query := "select customer_pass_code,sum(total) from order_list group by customer_pass_code;"
	passcode := ""
	total := 0
	err := tx.QueryRow(query).Scan(&passcode, &total)
	Validate(err, "select", tx)
	return total, passcode
}

// =======================================================Insert Report Detail Laundry ==============================================
func (r trxOrderListRepo) InsertReportDetail(report *models.DetailReportReq, tx *sql.Tx) string {
	query := "insert into report_detail (pass_code,paid_status,date_in,total_price) values($1,$2,$3,$4)"
	_, err := tx.Exec(query, &report.Pascode, &report.PaidStatus, time.Now(), &report.TotalPrice)
	Validate(err, "insert", tx)
	return fmt.Sprintf("Total %d", report.TotalPrice)
}
func (r trxOrderListRepo) TruncateOrder(tx *sql.Tx) {
	query := "Truncate table order_list"
	_, err := tx.Exec(query)
	Validate(err, "truncate", tx)

}

//============================================================To Order List================================================
//============================================================Insert order Laundry ========================================

func (r trxOrderListRepo) InsertOrder(order *models.OrderList, tx *sql.Tx) {
	query := "insert into order_list (id, customer_pass_code,service_id,quantity) values($1,$2,$3,$4)"

	_, err := tx.Exec(query, &order.ID, &order.CustomerPasscode, &order.ServiceId, &order.Quantity)
	Validate(err, "Insert", tx)

}

// ============================================================Get Report Order Data From orderList =====================================
func (r trxOrderListRepo) GetReport(id string, tx *sql.Tx) models.OrderReport {
	var reportOrder models.OrderReport

	query := `
select ol.id, c.name,c.pass_code,s.service,ol.quantity,s.unit, total from order_list as ol
join customer as c on ol.customer_pass_code =c.pass_code
join service as s on ol.service_id = s.id
where ol.id = $1
group by ol.id,c.name,c.pass_code,s.service ,ol.quantity,s.unit, ol.total;`

	err := tx.QueryRow(query, id).Scan(
		&reportOrder.OrderListId,
		&reportOrder.Name,
		&reportOrder.PassCode,
		&reportOrder.Service,
		&reportOrder.Quantity,
		&reportOrder.Unit,
		&reportOrder.Total,
	)
	Validate(err, "select", tx)
	return reportOrder
}

// ============================================================Insert Report From Get REport ==============================================
func (r trxOrderListRepo) InsertReportOrder(reportOrder *models.OrderReport, tx *sql.Tx) {
	query := "INSERT INTO report_order(order_list_id,name,pass_code,service,quantity,unit,total,created_at) values ($1,$2,$3,$4,$5,$6,$7,$8)"

	_, err := tx.Exec(query,
		&reportOrder.OrderListId,
		&reportOrder.Name,
		&reportOrder.PassCode,
		&reportOrder.Service,
		&reportOrder.Quantity,
		&reportOrder.Unit,
		&reportOrder.Total,
		time.Now(),
	)
	Validate(err, "insert", tx)

}

// ==========================================================Respot data order ===================================================================
func (r trxOrderListRepo) GetAllOrder(tx *sql.Tx) any {
	var orders []models.OrderList
	query := "select ol.id,ol.customer_pass_code,s.service,ol.quantity,ol.total from order_list as ol join service as s on ol.service_id = s.id"

	rows, err := tx.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.OrderList
		if err := rows.Scan(&order.ID, &order.CustomerPasscode, &order.ServiceId, &order.Quantity, &order.Total); err != nil {
			log.Println(err)
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		return nil
	}

	return orders
}

// =============================================================Update column Total from quantity * service.price========================================================
func (r trxOrderListRepo) UpdateTotalOrderList(tx *sql.Tx) {
	query := `update order_list set total = order_list.quantity * service.price
	from service
	where order_list.service_id = service.id;
	`
	_, err := tx.Exec(query)
	Validate(err, "Update", tx)

}

// =====================================Validate ====================================================
func Validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println("successfully " + message + " data")
	}
}
