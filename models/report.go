package models

type OrderReport struct {
	OrderListId string  `json:"order_list_id"`
	Name        string  `json:"name"`
	PassCode    string  `json:"pass_code"`
	Service     string  `json:"service"`
	Quantity    float32 `json:"quantity"`
	Unit        string  `json:"unit"`
	Total       int     `json:"total"`
	Created_At  string  `json:"created_at"`
}

type DetailReportReq struct {
	Pascode    string `json:"pass_code"`
	PaidStatus bool   `json:"paid_status"`
	DateIn     string `json:"date_in"`
	DateOut    string `json:"date_out"`
	TakenBy    string `json:"taken_by"`
	TotalPrice int    `json:"total_price"`
}

//=================================
