package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Ok      bool              `json:"ok"`
	Message string            `json:"message,omitempty"`
	Data    any               `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type OKResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Errors  map[string]string
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func OK(w http.ResponseWriter, r *http.Request, okResponse *OKResponse) {
	code := okResponse.Code
	if code == 0 {
		code = http.StatusOK
	}

	render.Status(r, code)
	render.Render(w, r, &Response{Ok: true, Message: okResponse.Message, Data: okResponse.Data})
}

func Error(w http.ResponseWriter, r *http.Request, errResponse *ErrResponse) {
	code := errResponse.Code
	if code == 0 {
		code = http.StatusInternalServerError
	}

	render.Status(r, code)
	render.Render(w, r, &Response{Ok: false, Message: errResponse.Message, Error: errResponse.Error, Errors: errResponse.Errors})
}
