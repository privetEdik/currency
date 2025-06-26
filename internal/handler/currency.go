package handler

import (
	"currency/internal/model"
	"currency/internal/repository"
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Error getting currencies", http.StatusInternalServerError)
		return
	}
	writeJSON(w, currencies)
}
func (h *Handler) GetByCode(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/currency/")
	if len(code) != 3 {
		http.Error(w, "Invalid currency code", http.StatusBadRequest)
	}
	currency, err := h.repo.GetByCode(strings.ToUpper(code))
	if err != nil {
		http.Error(w, "Error fitching  currency", http.StatusInternalServerError)
		return
	}
	if currency == nil {
		http.Error(w, "Currency not found", http.StatusNotFound)
	}
	writeJSON(w, currency)
}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, `{"message":"Invalid body"}`, http.StatusBadRequest)
		return
	}

	code := strings.ToUpper(r.FormValue("code"))
	name := r.FormValue("name")
	sign := r.FormValue("sign")
	if len(code) == 0 || len(name) == 0 || len(sign) == 0 {
		http.Error(w, `{"message":"Missing or invalid fields"}`, http.StatusBadRequest)
		return
	}

	c := model.Currency{
		Code: code,
		Name: name,
		Sign: sign,
	}
	if err := h.repo.Insert(c); err != nil {
		http.Error(w, `{"message":"Failed to insert currency"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	writeJSON(w, c)
}
func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
