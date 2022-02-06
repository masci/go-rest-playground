package main

import (
	"net/http"

	"github.com/go-chi/render"
)

/*
	Error types

	Our API will send back meaningful errors in JSON format.
*/

// ErrResponse is our Error type for HTTP responses
type ErrResponse struct {
	Err            error  `json:"-"`               // low-level runtime error
	HTTPStatusCode int    `json:"-"`               // http response status code
	StatusText     string `json:"status"`          // user-level status message
	ErrorText      string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render provides a representation of the error, same as valid responses
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrRender is returned when we weren't able to provide a representation of the
// resource requested
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// ErrInvalidRequest is for bad requests
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound is the classic 404
func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: 404,
		StatusText:     "Resource not found.",
	}
}
