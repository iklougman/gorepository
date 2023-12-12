package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"igor.dev/go/todo-with-handler/model"
)

type BaseHandler struct {
	todoRepo model.TodoRepo
}

func NewTodoHandler(todoRepo *model.TodoRepo) *BaseHandler {
	return &BaseHandler{
		todoRepo: *todoRepo,
	}
}

func (h *BaseHandler) MuxHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(os.Getenv("ENDPOINT"), h.Routes)

	return mux
}

func (h *BaseHandler) Routes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodPut:
	}

}

func (h *BaseHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoRepo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(todos)

}

func (h *BaseHandler) handlePost(w http.ResponseWriter, r *http.Request) {

}
