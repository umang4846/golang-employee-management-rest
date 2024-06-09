package main

import (
	"fmt"
	"golang-employee-management-rest/api"
	"golang-employee-management-rest/employee"
	"log"
	"net/http"
)

func main() {
	// Initialize employee store
	store := employee.NewStore()

	// Create some sample employees
	store.Create(employee.Employee{ID: 1, Name: "John Doe", Position: "Software Engineer", Salary: 75000})
	store.Create(employee.Employee{ID: 2, Name: "Jane Smith", Position: "Product Manager", Salary: 90000})
	store.Create(employee.Employee{ID: 3, Name: "Michael Johnson", Position: "Data Scientist", Salary: 85000})

	// Routes
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Request URL: GET /employees?page=<page_number>&pageSize=<page_size>
			api.ListEmployeesHandler(w, r, store)
		case http.MethodPost:
			// Request URL: POST /employees
			api.CreateEmployeeHandler(w, r, store)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Request URL: GET /employees/<employee_id>
			api.GetEmployeeHandler(w, r, store)
		case http.MethodPut:
			// Request URL: PUT /employees/<employee_id>
			api.UpdateEmployeeHandler(w, r, store)
		case http.MethodDelete:
			// Request URL: DELETE /employees/<employee_id>
			api.DeleteEmployeeHandler(w, r, store)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
