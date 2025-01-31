package server

import (
	"encoding/json"
	"net/http"
)


func BadRequest(slug string, w http.ResponseWriter, r *http.Request) {
	httpRespondWithErr(slug, w, r, "bad request", http.StatusBadRequest)
}

func InternalServer(slug string, w http.ResponseWriter, r *http.Request) {
	httpRespondWithErr(slug, w, r, "internal server", http.StatusInternalServerError)
}

func httpRespondWithErr(slug string, w http.ResponseWriter, r *http.Request, msg string, status int) {

	resp := ErrResponse{Slug: slug, Msg: msg, httpStatus: status}

	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}


type ErrResponse struct {
	Slug string `json:"slug"`
	Msg string `json:"msg"`
	httpStatus int
}