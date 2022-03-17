package models

import "time"

type Merchant struct {
	ID           int32     `json: "id"`
	UserId       int32     `json: "user_id"`
	MerchantName string    `json: "merchant_name"`
	CreatedAt    time.Time `json: "created_at"`
	CreatedBy    int32     `json: "created_by"`
	UpdatedAt    time.Time `json: "updated_at"`
	UpdatedBy    int32     `json: "updated_by"`
}
