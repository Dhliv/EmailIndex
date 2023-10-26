package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dhliv/EmailServer/src/requests"
	"github.com/go-chi/chi/v5"
)

func returnError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func GetEmails(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		returnError(w, fmt.Errorf("Error while retrieving param 'page': %v\n", err))
	}

	emails, err := requests.GetEmails(page)
	if err != nil {
		returnError(w, err)
		return
	}

	resp, err := json.Marshal(emails)
	if err != nil {
		returnError(w, err)
		return
	}
	
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}