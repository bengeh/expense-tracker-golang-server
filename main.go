package main

// import "expense_tracker_golang_api_with_docker/models"
import "expense_tracker_golang_api_with_docker/handlers"
import (
	// "context"
	"fmt"
	"log"

	"net/http"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Rest of the code will go here
	// Set client options
	router := mux.NewRouter()
	router.HandleFunc("/createExpenses", handlers.CreateExpenseEndPoint).Methods("POST")
	router.HandleFunc("/getExpenses", handlers.GetExpensesEndPoint).Methods("GET")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
	
}