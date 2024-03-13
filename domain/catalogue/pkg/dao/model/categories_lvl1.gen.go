// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCategoriesLvl1 = "categories_lvl1"

// CategoriesLvl1 mapped from table <categories_lvl1>
type CategoriesLvl1 struct {
	Slug        string    `gorm:"column:slug;primaryKey" json:"slug"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Status      int32     `gorm:"column:status;not null" json:"status"`
	CreatedDate time.Time `gorm:"column:created_date;not null;default:now()" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updated_date"`
	CreatedBy   string    `gorm:"column:created_by;not null;default:'system_default'::character varying" json:"created_by"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by"`
	IsVisible   bool      `gorm:"column:is_visible;not null;default:true" json:"is_visible"`
}

// TableName CategoriesLvl1's table name
func (*CategoriesLvl1) TableName() string {
	return TableNameCategoriesLvl1
}