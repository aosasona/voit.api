package handler

import (
	"net/http"

	"voit.api/pkg/response"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	response.OK(w, r, &response.OKResponse{Message: "pong"})
}
