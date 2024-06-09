package employee

import (
	"reflect"
	"sort"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	store := NewStore()

	// Create an employee
	emp := Employee{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   75000,
	}
	store.Create(emp)

	// Retrieve the employee
	retrievedEmp, err := store.GetByID(emp.ID)
	if err != nil {
		t.Errorf("Failed to retrieve employee: %v", err)
	}
	if retrievedEmp != emp {
		t.Errorf("Retrieved employee does not match the created one. Expected: %v, Got: %v", emp, retrievedEmp)
	}
}

func TestUpdateEmployee(t *testing.T) {
	store := NewStore()

	// Create an employee
	emp := Employee{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   75000,
	}
	store.Create(emp)

	// Update the employee
	updatedEmp := Employee{
		ID:       1,
		Name:     "John Doe Updated",
		Position: "Software Engineer",
		Salary:   80000,
	}
	err := store.Update(emp.ID, updatedEmp)
	if err != nil {
		t.Errorf("Failed to update employee: %v", err)
	}

	// Retrieve the updated employee
	retrievedEmp, err := store.GetByID(emp.ID)
	if err != nil {
		t.Errorf("Failed to retrieve updated employee: %v", err)
	}
	if retrievedEmp != updatedEmp {
		t.Errorf("Updated employee details do not match. Expected: %v, Got: %v", updatedEmp, retrievedEmp)
	}
}

func TestDeleteEmployee(t *testing.T) {
	store := NewStore()

	// Create an employee
	emp := Employee{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   75000,
	}
	store.Create(emp)

	// Delete the employee
	err := store.Delete(emp.ID)
	if err != nil {
		t.Errorf("Failed to delete employee: %v", err)
	}

	// Attempt to retrieve the deleted employee
	_, err = store.GetByID(emp.ID)
	if err == nil {
		t.Errorf("Employee was not deleted successfully")
	}
}

func TestGetEmployeeByIDNotFound(t *testing.T) {
	store := NewStore()

	// Attempt to retrieve a non-existent employee
	_, err := store.GetByID(100)
	if err == nil {
		t.Error("Expected error when retrieving non-existent employee, but got none")
	}
}

func TestUpdateEmployeeNotFound(t *testing.T) {
	store := NewStore()

	// Attempt to update a non-existent employee
	emp := Employee{
		ID:       100,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   75000,
	}
	err := store.Update(emp.ID, emp)
	if err == nil {
		t.Error("Expected error when updating non-existent employee, but got none")
	}
}

func TestDeleteEmployeeNotFound(t *testing.T) {
	store := NewStore()

	// Attempt to delete a non-existent employee
	err := store.Delete(100)
	if err == nil {
		t.Error("Expected error when deleting non-existent employee, but got none")
	}
}

func TestList(t *testing.T) {
	employees := map[int]Employee{
		1: {ID: 1, Name: "John Doe", Position: "Engineer", Salary: 50000},
		2: {ID: 2, Name: "Jane Smith", Position: "Manager", Salary: 60000},
		3: {ID: 3, Name: "Alice Johnson", Position: "Analyst", Salary: 45000},
	}

	store := &Store{
		employees: employees,
	}

	tests := []struct {
		page     int
		pageSize int
		expected []Employee
	}{
		{page: 1, pageSize: 2, expected: []Employee{employees[1], employees[2]}},
		{page: 2, pageSize: 2, expected: []Employee{employees[3]}},
		{page: 1, pageSize: 3, expected: []Employee{employees[1], employees[2], employees[3]}},
		{page: 2, pageSize: 3, expected: []Employee{employees[1], employees[2], employees[3]}},
		{page: 0, pageSize: 2, expected: []Employee{}},
		{page: 1, pageSize: 0, expected: []Employee{}},
	}

	for _, test := range tests {
		result, err := store.List(test.page, test.pageSize)
		if err != nil {
			t.Errorf("List failed with error: %v", err)
		}

		// Sort the result by ID for consistent comparison
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("List for page=%d, pageSize=%d: expected %v, got %v", test.page, test.pageSize, test.expected, result)
		}
	}
}
