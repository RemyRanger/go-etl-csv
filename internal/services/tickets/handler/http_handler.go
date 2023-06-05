package handler

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/RemyRanger/go-etl-csv/internal/common/http_error"
	"github.com/RemyRanger/go-etl-csv/internal/common/models"
	"github.com/RemyRanger/go-etl-csv/internal/services/tickets"
	"github.com/go-chi/render"
	"gopkg.in/yaml.v2"
)

const TicketHeadersMaxRow int = 4
const ProductHeadersRow int = 5

// HTTPHandler : handler http type
type HTTPHandler struct {
	svc tickets.Service
}

// NewHTTPHandler : create new http handler
func NewHTTPHandler(svc tickets.Service) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

// Create Ticket
// (POST /ticket)
func (h *HTTPHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()                // close the reader at the end of the function
	scanner := bufio.NewScanner(r.Body) // open scanner for incoming body

	ticket := &models.Ticket{}        // ticket sql model to be save in db
	headerMap := make(map[string]int) // to map column name and column index

	rowCount := 0        // keep row count
	for scanner.Scan() { // scan row by row
		rowCount++

		// Handle ticket headers delimited by row number
		if rowCount <= TicketHeadersMaxRow {
			text := scanner.Text()
			fmt.Println(text)
			if err := yaml.Unmarshal(scanner.Bytes(), ticket); err != nil {
				render.Render(w, r, http_error.ErrInternal(err, http_error.ErrInternalProcess))
			}

			continue
		}

		// New Csv reader for this row
		csvReader := csv.NewReader(bytes.NewReader(scanner.Bytes()))
		// Read row as csv
		record, err := csvReader.Read()
		if err != nil {
			render.Render(w, r, http_error.ErrInternal(err, http_error.ErrInternalProcess))
		}

		// Handle csv headers in one row, column doesn't have to be orderer, changing column doesn't affect the process
		if rowCount == ProductHeadersRow {
			// Mapping column name and index
			for i, v := range record {
				headerMap[v] = i
			}

			continue
		}

		// Handle CSV datas : if more column are added to the csv format we just have to add the new field here
		ticket.Products = append(ticket.Products, &models.Product{
			Name:      &record[headerMap["product"]],
			ProductId: &record[headerMap["product_id"]],
			Price:     &record[headerMap["price"]],
		})
	}

	err := h.svc.CreateTicket(r.Context(), ticket) // create ticket to db with service and next repository
	if err != nil {
		render.Render(w, r, http_error.ErrInternal(err, http_error.ErrInternalProcess))
		return
	}

	render.Status(r, http.StatusCreated)
}
