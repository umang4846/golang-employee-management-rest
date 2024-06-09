package api

import (
	"encoding/json"
	"golang-employee-management-rest/employee"
	"net/http"
	"strconv"
)

// ListEmployeesHandler handles the request to list employees with pagination.
func ListEmployeesHandler(w http.ResponseWriter, r *http.Request, store *employee.Store) {

	// Pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	// Retrieve employees from the store
	employees, err := store.List(page, pageSize)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Response
	json.NewEncoder(w).Encode(employees)
}

// CreateEmployeeHandler handles the request to create a new employee.
func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request, store *employee.Store) {
	// Parse request body
	var emp employee.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add employee to store
	store.Create(emp)
	w.WriteHeader(http.StatusCreated)
}

// GetEmployeeHandler handles the request to get an employee by ID.
func GetEmployeeHandler(w http.ResponseWriter, r *http.Request, store *employee.Store) {
	// Extract employee ID from request URL
	id, err := strconv.Atoi(r.URL.Path[len("/employees/"):])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Retrieve employee from store
	emp, err := store.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Respond with employee details
	json.NewEncoder(w).Encode(emp)
}

// UpdateEmployeeHandler handles the request to update an existing employee.
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request, store *employee.Store) {
	// Extract employee ID from request URL
	id, err := strconv.Atoi(r.URL.Path[len("/employees/"):])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var emp employee.Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update employee in store
	err = store.Update(id, emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteEmployeeHandler handles the request to delete an employee by ID.
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request, store *employee.Store) {
	// Extract employee ID from request URL
	id, err := strconv.Atoi(r.URL.Path[len("/employees/"):])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Delete employee from store
	err = store.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
