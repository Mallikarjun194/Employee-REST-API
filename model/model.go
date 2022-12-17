package model

type Employee struct {
	EmpID   string  `json:"emp_id"`
	EmpName string  `json:"emp_name,omitempty"`
	Salary  float32 `json:"salary,omitempty"`
	Age     int     `json:"age,omitempty"`
}

type Response struct {
	EmpID   string `json:"emp_id,omitempty"`
	Message string `json:"message,omitempty"`
}