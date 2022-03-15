package models

type Transaction struct {
	ID         int32   `json: "id"`
	MerchantID int32   `json: "merchant_id"`
	OutletID   int32   `json: "outlet_id"`
	BillTotal  float32 `json: "bill_total"`
}
