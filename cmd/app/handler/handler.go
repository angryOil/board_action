package handler

import (
	"board_action/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
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

}

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {

}
