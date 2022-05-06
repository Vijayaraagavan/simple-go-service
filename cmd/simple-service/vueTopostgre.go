package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)

type numbers struct {
	Num1 float64 `json:"num"`
	Num2 string `json:"num2"`
}

type numsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
}

func process(numsdata numbers) (numsResponseData) {
	
	var numsres numsResponseData
	fl,_:=  strconv.ParseFloat(numsdata.Num2,64)
	numsres.Add = numsdata.Num1 + fl
	numsres.Mul = numsdata.Num1 * fl
	// numsres.Sub = numsdata.Num1 - numsdata.Num2
	// numsres.Div = numsdata.Num1 / numsdata.Num2

	// Postgresql example
	// connecting database
	db,err := sql.Open("postgres","user=postgres password=ard5050 host=127.0.0.1 port=5432 dbname=vijay sslmode=disable")
	if err!=nil {
		panic(err)
	} else {
		fmt.Println("connected successfully")
	}

	// testing the connection
	defer db.Close()
	connectivity := db.Ping()
	if connectivity!=nil {
		panic(err)
	} else {
		fmt.Println("conection test is good")
	}

/*	//create a table
	Dbcreate := `
	CREATE TABLE teams(
		id integer,
		team_names varchar(100)
	);`
	_, err = db.Exec(Dbcreate)
	if err!=nil {
		panic(err)
	} else {
		fmt.Println("New table is created")
	} */

	//Prepare a query
	insert,err := db.Prepare("INSERT INTO teams VALUES ($1,$2)")
	if err!=nil {
		panic(err)
	} else {
		fmt.Println("Prepared a insert query")
	}

	//Insert values into table using prepared statement
	_,err = insert.Exec(numsdata.Num1,numsdata.Num2)
	if err!=nil {
		panic(err)
	} else {
		fmt.Println("A row is inserted successfully")
	}
	insert.Close() 

	return numsres
}

func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numsData numbers
	var numsResData numsResponseData
	
	decoder.Decode(&numsData)

	numsResData = process(numsData)

	fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(numsResData); err != nil {
        panic(err)
    }
}

func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}