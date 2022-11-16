package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var employees []Employee

func values() {
	employees = append(employees, Employee{Id: "1", Name: "Ahin", Location: "Thottavaram"})
	employees = append(employees, Employee{Id: "2", Name: "Vinith", Location: "Moovathumugam"})
}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range employees {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	employees = append(employees, emp)
	w.Write([]byte("Success"))
}
func main() {
	r := mux.NewRouter()
	values()
	r.HandleFunc("/employees", GetEmployee).Methods("GET")
	r.HandleFunc("/employee/{id}", GetById).Methods("GET")
	r.HandleFunc("/employee", AddEmployee).Methods("POST")
	fmt.Println("server running at 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
