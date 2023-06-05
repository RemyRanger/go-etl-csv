package models

import (
	"time"
)

// Product : database Product entity model
type Product struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        int64   `gorm:"primary_key;type:int"`
	TicketId  int64   `gorm:"column:ticket_id;type:int"`
	Name      *string `gorm:"column:name;type:varchar"` // pointer on string for possible null field, non-blocking if empty message
	ProductId *string `gorm:"column:product_id;type:varchar"`
	Price     *string `gorm:"column:price;type:varchar"`
	Ticket    Ticket
}

// TableName sets the insert table name for this struct type
func (p *Product) TableName() string {
	return "products"
}
