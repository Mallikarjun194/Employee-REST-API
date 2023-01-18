package handlers

import (
	"Employee-REST-API/model"
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

var EmpMap = make(map[string]*model.Employee, 0)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	NewEmp := &model.Employee{}
	var empObj []byte

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, NewEmp)

	EmpMap[NewEmp.EmpID] = NewEmp
	empObj, _ = json.Marshal(NewEmp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(empObj)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var Result []*model.Employee

	for _, value := range EmpMap {
		Result = append(Result, value)
	}
	empObj, _ = json.Marshal(Result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var empID = chi.URLParam(r, "empID")

	empDetail, _ := EmpMap[empID]
	empObj, _ = json.Marshal(empDetail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var empID = chi.URLParam(r, "empID")

	empDetail, _ := EmpMap[empID]
	body, _ := ioutil.ReadAll(r.Body)

	emp := &model.Employee{}
	json.Unmarshal(body, emp)

	EmpMap[empID].Age = emp.Age
	EmpMap[empID].Salary = emp.Salary
	empObj, _ = json.Marshal(empDetail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var empID = chi.URLParam(r, "empID")

	delete(EmpMap, empID)
	empObj, _ = json.Marshal("Successfully deleted the emp " + empID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)
}
