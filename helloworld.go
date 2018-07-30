package main

import (
    "log"
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/cpwr/go_test/database"
)

type Person struct {
    ID int `json:"id,omitempty"`
    Name string `json:"firstname,omitempty"`
    Created string `json:"lastname,omitempty"`
}

type Task struct {
    ID int `json:"id,omitempty"`
    Taskname string `json:"taskname,omitempty"`
}

var tasks []Task
var persons []Person
var db *sql.DB

func Query() error {
    name := "Aaron"
    rows, err := db.Query("SELECT id, name, created FROM employees where name=$1", name)
    if err != nil {
        return err
    }

    for rows.Next() {
        var p Person
        if err := rows.Scan(&p.ID, &p.Name, &p.Created); err != nil {
            return err
        }
        persons = append(persons, p)
    }
    return rows.Err()
}

func GetInfo(w http.ResponseWriter, req *http.Request) {
    Query()
    json.NewEncoder(w).Encode(persons)
}

func UpdateInfo(w http.ResponseWriter, req *http.Request) {

}

func DeleteInfo(w http.ResponseWriter, req *http.Request) {

}

func HomePage(w http.ResponseWriter, req *http.Request) {
    tasks = append(tasks, Task{ID: 1, Taskname: "Estimation"})
    tasks = append(tasks, Task{ID: 2, Taskname: "Estimate estimation"})
    json.NewEncoder(w).Encode(tasks)
}


func requestHandler() {
    router := mux.NewRouter()
    router.HandleFunc("/", HomePage)
    router.HandleFunc("/users/{name}", GetInfo).Methods("GET")
    router.HandleFunc("/users/{name}", UpdateInfo).Methods("PUT")
    router.HandleFunc("/users/{name}", DeleteInfo).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8001", router))
}

func main() {
    var err error
    if db, err = database.CreateConn(db); err != nil {
        panic(err)
    }
    defer db.Close()

    if err = database.CreateTable(db); err != nil {
        panic(err)
    }
    requestHandler()
}
