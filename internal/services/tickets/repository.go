package tickets

import (
	"context"

	"github.com/RemyRanger/go-etl-csv/internal/common/models"
)

// Repository : interface for repository
type Repository interface {
	CreateTicket(context.Context, *models.Ticket) error
}
