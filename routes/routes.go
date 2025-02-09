package routes

import (
	"go-server/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos/add", handlers.AddTodo).Methods("POST")
	router.HandleFunc("/todos/{id}/update", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}/delete", handlers.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

	// http.HandleFunc("/todos", handlers.GetTodos)
	// http.HandleFunc("/todos/add", handlers.AddTodo)
}
