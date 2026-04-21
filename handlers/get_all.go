package handlers

import (
	"api-todo-aula2/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Ocorreu um erro ao buscar os registros: %s ->", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
