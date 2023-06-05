package tickets

import (
	"context"

	"github.com/RemyRanger/go-etl-csv/internal/common/logs"
	"github.com/RemyRanger/go-etl-csv/internal/common/models"
)

// ServiceName : service name
const ServiceName string = "tickets"

// Service : Service interface
type Service interface {
	CreateTicket(context.Context, *models.Ticket) error
}

type service struct {
	repository Repository
}

// NewService : create service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// CreateTicket : create a ticket
func (srv *service) CreateTicket(ctx context.Context, ticket *models.Ticket) error {
	if err := srv.repository.CreateTicket(ctx, ticket); err != nil {
		logs.LogError(ServiceName, "Error create ticket in database.", err)
		return err
	}

	/* if err := srv.repository.CreateProducts(ctx, products); err != nil {
		logs.LogError(ServiceName, "Error create products in database.", err)
		return err
	} */

	return nil
}
