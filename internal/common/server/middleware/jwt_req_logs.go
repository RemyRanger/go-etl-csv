package middleware

import (
	"io"
	"net/http"

	"github.com/RemyRanger/go-etl-csv/internal/common/logs"
	"github.com/RemyRanger/go-etl-csv/internal/common/models"
	"gorm.io/gorm"
)

// ServiceRequestLogsMiddlewareName : service name
const ServiceRequestLogsMiddlewareName string = "jwt_ticket_logs"

// RequestLogsMiddleware : RequestLogsMiddleware
type RequestLogsMiddleware struct {
	db *gorm.DB
}

// NewRequestLogsMiddleware creates a new RequestLogsMiddleware handler with the provided options.
func NewRequestLogsMiddleware(db *gorm.DB) *RequestLogsMiddleware {
	return &RequestLogsMiddleware{
		db: db,
	}
}

// Handler is a middleware that logs in database incoming ticket
func (fc *RequestLogsMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(fc.db, r)

		next.ServeHTTP(w, r)
	})
}

// logRequest : middleware for logging each incoming request with payload to keep track.
func logRequest(db *gorm.DB, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	ticket := string(payload)

	if err != nil {
		logs.LogError(ServiceRequestLogsMiddlewareName, "Error extracting payload from request", err)
	}

	requestLog := &models.RequestLog{
		Payload: &ticket,
	}

	if err := db.Create(&requestLog).Error; err != nil {
		logs.LogError(ServiceRequestLogsMiddlewareName, "Error creating request logs in db", err)
	}
}
