package handlers

import (
	"fmt"
	"encoding/json"
	"expense_tracker_golang_api_with_docker/dao"
	"expense_tracker_golang_api_with_docker/models"
	"net/http"
)


// GetAllExpenseEndpoint gets all expenses
func CreateExpenseEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HENLO")
	
	name := r.FormValue("name")
	price := r.FormValue("price")
	newExpense := models.Expense{
		Name: name,
		Price: price,
	}
	
	dao.InsertOneValue(newExpense)
	json.NewEncoder(w).Encode(newExpense)
}

// // GetExpenseEndpoint gets 1 expense
func GetExpensesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BYEBYE FROM ENDPOINT GET")
	params := r.Form
	payload := dao.GetAllExpense()
	fmt.Println(payload)
	for _, p := range payload {
		if p.Name == params.Get("name") {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Person not found")
}

// // CreatePersonEndpoint creta a person
// func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
// 	var person models.Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	dao.InsertOneValue(person)
// 	json.NewEncoder(w).Encode(person)
// }

// // DeletePersonEndpoint delets a person
// func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
// 	var person models.Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	dao.DeletePerson(person)
// }

// // UpdatePersonEndpoint updates a person
// func UpdatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
// 	personID := mux.Vars(r)["id"]
// 	var person models.Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	dao.UpdatePerson(person, personID)

// }
