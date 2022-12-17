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
	var Resp model.Response
	var empObj []byte

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Resp.Message = fmt.Sprintf("Error while reading request: %s", err)
		empObj, _ = json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(empObj)
		if err != nil {
			return
		}
		return
	}
	err = json.Unmarshal(body, NewEmp)
	if err != nil {
		Resp.Message = fmt.Sprintf("Error while unmarshelling: %s", err)
		empObj, _ = json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // which status should we mention
		w.Write(empObj)

	}

	if _, ok := EmpMap[NewEmp.EmpID]; ok {
		Resp.Message = fmt.Sprintf("Emp %s details already present ", NewEmp.EmpID)
		empObj, _ = json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(empObj)
		return
	}

	EmpMap[NewEmp.EmpID] = NewEmp

	Resp.EmpID = NewEmp.EmpID
	Resp.Message = fmt.Sprintf("Successfully added emp with ID %s ", NewEmp.EmpID)

	empObj, _ = json.Marshal(Resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(empObj)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var Resp model.Response
	if len(EmpMap) < 1 {
		Resp.Message = "Employees details not present"
		empObj, _ := json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // which status should we mention
		w.Write(empObj)
		return
	}
	var Result []*model.Employee
	for _, value := range EmpMap {
		Result = append(Result, value)
	}
	empObj, _ = json.Marshal(Result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // which status should we mention
	w.Write(empObj)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var Resp model.Response
	var empID = r.URL.Query().Get("empID")

	empDetail, ok := EmpMap[empID]
	if !ok {
		Resp.Message = fmt.Sprintf("Emp %s details not present", empID)
		empObj, _ := json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // which status should we mention
		w.Write(empObj)
		return
	}
	empObj, _ = json.Marshal(empDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // which status should we mention
	w.Write(empObj)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var Resp model.Response
	var empID = r.URL.Query().Get("empID")

	empDetail, ok := EmpMap[empID]
	if !ok {
		Resp.Message = fmt.Sprintf("Emp %s details not present", empID)
		empObj, _ = json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // which status should we mention
		w.Write(empObj)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Resp.Message = fmt.Sprintf("Error while reading request: %s", err)
		empObj, _ = json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(empObj)
		if err != nil {
			return
		}
		return
	}
	emp := &model.Employee{}
	_ = json.Unmarshal(body, emp)
	EmpMap[empID].Age = emp.Age
	EmpMap[empID].Salary = emp.Salary
	empObj, _ = json.Marshal(empDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // which status should we mention
	w.Write(empObj)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	var empObj []byte
	var Resp model.Response
	var empID = r.URL.Query().Get("empID")
	_, ok := EmpMap[empID]
	if !ok {
		Resp.Message = fmt.Sprintf("Emp %s details not present", empID)
		empObj, _ = json.Marshal(Resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // which status should we mention
		w.Write(empObj)
		return
	}
	delete(EmpMap, empID)
	Resp.Message = "Successfully deleted the emp " + empID
	empObj, _ = json.Marshal(Resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound) // which status should we mention
	w.Write(empObj)
}
