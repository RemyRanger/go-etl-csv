package repository

import (
	"context"

	"github.com/RemyRanger/go-etl-csv/internal/common/models"
)

// CreateTicket : create ticket in database
func (r *SQLRepository) CreateTicket(ctx context.Context, ticket *models.Ticket) error {
	if err := r.DB.Create(&ticket).Error; err != nil {
		return err
	}

	return nil
}
