# Employee-REST-API
## CRUD API's for Employe Details

### Clone the repo:
    git clone https://github.com/Mallikarjun194/Employee-REST-API.git
####
#### To Start the server:
    go run main.go
####
### API END-POINTS:
###
##### 1.  CREATE EMP: Create empolyee with few details
    EndPoint:  http://localhost:5001/emp 

Payload:

    {
        "emp_id":"1234567",
        "emp_name":"Arjun",
        "salary":76999,
        "age":29
    }

Response:

    {
        "emp_id": "1234567",
        "emp_name": "Arjun",
        "salary": 76999,
        "age": 29
    }

###

##### 2. GET All Employees: GET call to fetch all the emp details
    EndPoint: http://localhost:5001/emp

Response:

    [
        {
            "emp_id": "1234567",
            "emp_name": "Arjun",
            "salary": 76999,
            "age": 29
        },
        {
            "emp_id": "123456",
            "emp_name": "AJ",
            "salary": 7699,
            "age": 25
        }
    ]

###
##### 3. GETByID: GET Call by ID to fetch details of a particular employee
    EndPoint: http://localhost:5001/emp/1234567

Response:

    {
        "emp_id": "1234567",
        "emp_name": "Arjun",
        "salary": 76999,
        "age": 29
    }

 #
##### 4. UPDATE: Update by ID to update particular employee details
    EndPoint: http://localhost:5001/emp/1234567

Payload:

    {
        "salary":678,
        "age":39
    }

Response:

    {
        "emp_id": "1234567",
        "emp_name": "Arjun",
        "salary": 678,
        "age": 39
    }

###
##### 5. DELETE: Delete by ID to delete particular employee details
    EndPoint: http://localhost:5001/emp/1234567

Response:

    "Successfully deleted the emp 1234567"




