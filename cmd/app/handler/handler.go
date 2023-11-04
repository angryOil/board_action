package handler

import (
	"board_action/internal/controller"
	"board_action/internal/controller/req"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}", h.getInfo).Methods(http.MethodGet)
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}", h.create).Methods(http.MethodPost)
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

func (h Handler) getInfo(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	boardTypeId, err := strconv.Atoi(vars["boardTypeId"])
	if err != nil {
		http.Error(w, "invalid board type id", http.StatusBadRequest)
		return
	}
	var d req.CreateDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Create(r.Context(), cafeId, boardTypeId, d)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {

}
