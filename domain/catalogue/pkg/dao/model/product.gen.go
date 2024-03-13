// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProduct = "product"

// Product mapped from table <product>
type Product struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	MerchantID  string    `gorm:"column:merchant_id;not null" json:"merchant_id"`
	Status      int32     `gorm:"column:status;not null" json:"status"`
	CreatedDate time.Time `gorm:"column:created_date;not null;default:now()" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updated_date"`
	CreatedBy   string    `gorm:"column:created_by;not null;default:'system_default'::character varying" json:"created_by"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}