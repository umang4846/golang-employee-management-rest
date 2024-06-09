package employee

import (
	"fmt"
	"sort"
	"sync"
)

// Employee struct represents an employee entity.
type Employee struct {
	ID       int     `json:"id,omitempty"` // Unique identifier for the employee.
	Name     string  // Name of the employee.
	Position string  // Position/title of the employee.
	Salary   float64 // Salary of the employee.
}

// Store manages the storage and operations for employees.
type Store struct {
	sync.RWMutex
	employees map[int]Employee // Map to store employees by their ID.
}

// NewStore creates and returns a new instance of Store.
func NewStore() *Store {
	return &Store{
		employees: make(map[int]Employee),
	}
}

// Create adds a new employee to the store.
func (s *Store) Create(emp Employee) {
	s.Lock()
	defer s.Unlock()
	s.employees[emp.ID] = emp
}

// GetByID retrieves an employee from the store by ID.
func (s *Store) GetByID(id int) (Employee, error) {
	s.RLock()
	defer s.RUnlock()
	emp, ok := s.employees[id]
	if !ok {
		return Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return emp, nil
}

// Update updates the details of an existing employee.
func (s *Store) Update(id int, emp Employee) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.employees[id]; !ok {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	s.employees[id] = emp
	return nil
}

// Delete removes an employee from the store by ID.
func (s *Store) Delete(id int) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.employees[id]; !ok {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	delete(s.employees, id)
	return nil
}

// List retrieves a list of employees from the in-memory store with pagination support.
func (s *Store) List(page, pageSize int) ([]Employee, error) {
	var employees []Employee
	for _, emp := range s.employees {
		employees = append(employees, emp)
	}

	// Sort employees by ID
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].ID < employees[j].ID
	})

	// Apply pagination
	start := (page - 1) * pageSize
	end := page * pageSize
	if start < 0 || start >= len(employees) {
		start = 0
	}
	if end > len(employees) {
		end = len(employees)
	}
	if start > len(employees) {
		start = len(employees)
	}
	if end < 0 {
		end = 0
	}
	return employees[start:end], nil
}
