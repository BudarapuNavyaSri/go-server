package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/models"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
)

var (
	todos = []models.Todo{
		{ID: "1", Task: "Learn Golang", Completed: false},
		{ID: "2", Task: "Build a Server", Completed: false},
	}
	mutex sync.Mutex
)

// GET
func GetTodos(w http.ResponseWriter, r *http.Request) {
	//logs the method (GET) and URL path (/todos)
	fmt.Printf("Received Request: %s %s\n", r.Method, r.URL.Path)

	fmt.Printf("current todos: %+v\n", todos)

	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
		fmt.Println("Sent response with status : 500 Internal Server Error")
		return
	}

	fmt.Println("Sent response with status: 200 OK")
}

// POST
func AddTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddTodo handler called")
	mutex.Lock()
	defer mutex.Unlock()

	//creating a new variable of type models.Todo. It refers to the struct defined in models package
	var newTodo models.Todo

	//json.NewDecoder(r.Body) creates a new JSON decoder that reads from r.Body
	//.Decode(&NewTodo) decodes the JSON data from request body into the newTodo variable
	//&newTodo is a pointer to newTodo variable, so the decoder can fill it with the parsed JSON data
	//The incoming JSON data (like {"ID": "3", "Task": "Learn Go", "Completed": false}) is decoded into a Go struct (models.Todo in this case).
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		fmt.Println("Failed to decode request body into newTodo, sending 400 Bad Request")
		return
	}


	fmt.Printf("Adding New Todo: %+v\n", newTodo)

	fmt.Println("Before appending:", todos)
	todos = append(todos, newTodo)
	fmt.Println("After appending:", todos)
	

	// todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// It encodes the Go struct (newTodo) into JSON format and sends it back as a response to the client.
	//When we pass a pointer to the encoder (like &newTodo), the encoder would modify the original data.
	//However, in this case, we don't need to modify the original struct; we just want to encode it into JSON and send it as a response.
	err = json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		http.Error(w, "failed to encode new Todo", http.StatusInternalServerError)
		fmt.Println("Error encoding new todo, sending 500 Internal server error")
		return
	}
	fmt.Println("Sent response with status: 201 created")
}

// update
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	//extract id from url
	id := mux.Vars(r)["id"]
	fmt.Printf("Updating Todo with ID: %s\n", id)

	//decode request body into updatedTodo struct
	var updatedTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		fmt.Println("Failed to decode request body into updatedTodo, sending 400 bad request")
		return
	}

	//locking the todos array to prevent concurrent modifications
	mutex.Lock()
	defer mutex.Unlock()

	//find and update the todo
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(todos[i])
			if err != nil {
				http.Error(w, "Failed to encode updated todo", http.StatusInternalServerError)
				fmt.Println("Error encoding updated todo, sending 500 Internal Server Error")
				return
			}
			fmt.Printf("updated Todo: %+v\n", todos[i])
			fmt.Println("Sent response with status: 200 OK")
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
	fmt.Println("Todo not found, sending 404 not found")
}

// delete
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Printf("Deleting Todo with ID: %s\n", id)

	mutex.Lock()
	defer mutex.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			fmt.Printf("Deleted Todo: %+v\n", todo)
			fmt.Println("Sent response with status: 204 No content")
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
	fmt.Println("Todo not found, sending 404 Not Found")
}
