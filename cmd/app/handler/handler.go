package handler

import (
	"board_action/internal/controller"
	"board_action/internal/controller/req"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
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
	// 조회
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}", h.getInfo).Methods(http.MethodGet)
	// 나머진 카페 주인만 가능함
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}", h.create).Methods(http.MethodPost)
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/board-actions/{cafeId:[0-9]+}/{boardTypeId:[0-9]+}/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InternalServerError = "internal server error"
	InvalidCafeId       = "invalid cafe id"
	InvalidBoardTypeId  = "invalid board type id"
	InvalidId           = "invalid id"
)

func (h Handler) getInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardTypeId, err := strconv.Atoi(vars["boardTypeId"])
	if err != nil {
		http.Error(w, InvalidBoardTypeId, http.StatusBadRequest)
		return
	}

	d, err := h.c.GetInfo(r.Context(), cafeId, boardTypeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(d)
	if err != nil {
		log.Println("getInfo json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardTypeId, err := strconv.Atoi(vars["boardTypeId"])
	if err != nil {
		http.Error(w, InvalidBoardTypeId, http.StatusBadRequest)
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardTypeId, err := strconv.Atoi(vars["boardTypeId"])
	if err != nil {
		http.Error(w, InvalidBoardTypeId, http.StatusBadRequest)
		return
	}
	var d req.PatchDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Patch(r.Context(), id, cafeId, boardTypeId, d)
	if err != nil {
		if strings.Contains(err.Error(), "no row") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardTypeId, err := strconv.Atoi(vars["boardTypeId"])
	if err != nil {
		http.Error(w, InvalidBoardTypeId, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	err = h.c.Delete(r.Context(), cafeId, boardTypeId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
