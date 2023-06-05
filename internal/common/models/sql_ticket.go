package models

import (
	"time"
)

// Ticket : database Ticket entity model
type Ticket struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        int64    `gorm:"primary_key;type:int"`
	Order     *int64   `yaml:"Order" gorm:"column:reference;type:int"`
	VAT       *float64 `yaml:"VAT" gorm:"column:vat;type:float"`
	Total     *float64 `yaml:"Total" gorm:"column:total;type:float"`
	Products  []*Product
}

// TableName sets the insert table name for this struct type
func (t *Ticket) TableName() string {
	return "tickets"
}
