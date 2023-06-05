package http_error

import (
	"net/http"

	"github.com/go-chi/render"
)

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error  `json:"-"`                      // low-level runtime error
	StatusText     string `json:"errorMessage,omitempty"` // user-level status message
	DetailsText    string `json:"details,omitempty"`      // detailed error message
	ErrorText      string `json:"-"`                      // application-level error message, for debugging
	AppCode        int64  `json:"-"`                      // application-specific error code
	HTTPStatusCode int    `json:"-"`                      // http response status code
}

const (
	// ErrorColor : red color for err logs
	ErrorColor = "\033[1;31m%s\033[0m"
)

// Render : render
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest : invalid request error
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     http.StatusText(http.StatusBadRequest),
		DetailsText:    err.Error(),
	}
}

// ErrDataConflict : conflict error
func ErrDataConflict(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		StatusText:     http.StatusText(http.StatusConflict),
		ErrorText:      message,
	}
}

// ErrDataNotFound : not found error
func ErrDataNotFound(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     http.StatusText(http.StatusNotFound),
		ErrorText:      message,
		DetailsText:    message,
	}
}

// ErrInternal : internal error
func ErrInternal(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     http.StatusText(http.StatusInternalServerError),
		ErrorText:      message,
	}
}

// ErrRender : render error
func ErrRender(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     http.StatusText(http.StatusUnprocessableEntity),
		ErrorText:      message,
	}
}

// Unauthorised : Unauthorised error
func Unauthorised(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     http.StatusText(http.StatusUnauthorized),
		ErrorText:      message,
	}
}

// Forbidden : Forbidden error
func Forbidden(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		StatusText:     http.StatusText(http.StatusForbidden),
		ErrorText:      message,
	}
}
