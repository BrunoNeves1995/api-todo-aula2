package apitodoaula2

import (
	"api-todo-aula2/configs"
	"api-todo-aula2/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.Get)
	r.Get("/", handlers.GetAll)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPI()), r)

}
