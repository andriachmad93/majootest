package models

import "time"

type Transaction struct {
	ID         int32     `json: "id"`
	MerchantID int32     `json: "merchant_id"`
	OutletID   int32     `json: "outlet_id"`
	BillTotal  float32   `json: "bill_total"`
	CreatedAt  time.Time `json: "created_at"`
	CreatedBy  int32     `json: "created_by"`
	UpdatedAt  time.Time `json: "updated_at"`
	UpdatedBy  int32     `json: "updated_by"`
}
