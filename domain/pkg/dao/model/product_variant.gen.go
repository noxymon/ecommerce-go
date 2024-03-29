// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProductVariant = "product_variant"

// ProductVariant mapped from table <product_variant>
type ProductVariant struct {
	ID                    string    `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name                  string    `gorm:"column:name;not null" json:"name"`
	Price                 float64   `gorm:"column:price;not null" json:"price"`
	SpecialPrice          float64   `gorm:"column:special_price" json:"special_price"`
	SpecialPriceStartDate time.Time `gorm:"column:special_price_start_date" json:"special_price_start_date"`
	SpecialPriceEndDate   time.Time `gorm:"column:special_price_end_date" json:"special_price_end_date"`
	StockAvailable        int32     `gorm:"column:stock_available;not null" json:"stock_available"`
	ProductID             int64     `gorm:"column:product_id;not null" json:"product_id"`
	Status                int32     `gorm:"column:status;not null" json:"status"`
	IsMainVariant         bool      `gorm:"column:is_main_variant;not null" json:"is_main_variant"`
	CreatedDate           time.Time `gorm:"column:created_date;not null;default:now()" json:"created_date"`
	UpdatedDate           time.Time `gorm:"column:updated_date" json:"updated_date"`
	CreatedBy             string    `gorm:"column:created_by;not null;default:'system_default'::character varying" json:"created_by"`
	UpdatedBy             string    `gorm:"column:updated_by" json:"updated_by"`
}

// TableName ProductVariant's table name
func (*ProductVariant) TableName() string {
	return TableNameProductVariant
}
