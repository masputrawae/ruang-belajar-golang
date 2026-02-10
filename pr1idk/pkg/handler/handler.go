package handler

import (
	"encoding/json"
	"net/http"
	"pr1idk/pkg/repository"
)

type HandlerManage struct {
	Repo repository.TodoManage
}

func New(r HandlerManage) *HandlerManage {
	return &r
}

func (hm *HandlerManage) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	results, err := hm.Repo.FindAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
