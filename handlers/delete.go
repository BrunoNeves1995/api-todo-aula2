package handlers

import (
	"api-todo-aula2/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Ocorreu um erro ao tentar fazer o parse do id: %s ->", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowDelete, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Ocorreu um erro ao deletar o registro: %s ->", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowDelete > 1 {
		log.Printf("Error: foram deletados %d registros", rowDelete)
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "Registro deletado com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
