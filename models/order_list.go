package models

type OrderList struct {
	ID               string  `json:"id"`
	CustomerPasscode string  `json:"customer_pass_code"`
	ServiceId        string  `json:"service_id"`
	Quantity         float32 `json:"qty"`
	Total            int     `json:"total"`
}

type OrderListReq struct {
	ID        string  `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Phone     string  `json:"phone" validate:"required"`
	ServiceId string  `json:"service_id"`
	Quantity  float32 `json:"qty"`
	Total     int     `json:"total"`
}
