package repository

import (
	"github.com/RemyRanger/go-etl-csv/internal/services/tickets"
	"gorm.io/gorm"
)

// SQLRepository : mysqlDB link
type SQLRepository struct {
	DB *gorm.DB
}

// NewSQLFlows : initializes a new FlowSqlRepo struct.
func NewRepository(database *gorm.DB) tickets.Repository {
	return &SQLRepository{
		DB: database,
	}
}
