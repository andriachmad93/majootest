package models

import "gorm.io/gorm"

func GetReport(db *gorm.DB, userid float64) ([]Transaction, error) {
	var report []Transaction

	result := db.Joins("JOIN merchants ON transactions.merchant_id = merchants.id and merchants.user_id = ? ", userid).Find(&report)

	return report, result.Error
}
