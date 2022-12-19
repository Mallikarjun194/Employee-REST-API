package controllers

import (
	"Employee-REST-API/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var EmpMap = make(map[string]*model.Employee, 0)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	NewEmp := &model.Employee{}
	Resp := &model.Response{}
	var empObj []byte

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, NewEmp)

	EmpMap[NewEmp.EmpID] = NewEmp
	Resp.Message = fmt.Sprintf("Successfully added emp with ID " + NewEmp.EmpID)
	empObj, _ = json.Marshal(Resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	var empID = r.URL.Query().Get("empID")

	empDetail, _ := EmpMap[empID]
	empObj, _ = json.Marshal(empDetail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var empID = r.URL.Query().Get("empID")

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
	var Resp model.Response
	var empID = r.URL.Query().Get("empID")

	delete(EmpMap, empID)
	Resp.Message = "Successfully deleted the emp " + empID
	empObj, _ = json.Marshal(Resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)
}
