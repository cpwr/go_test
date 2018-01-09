package main

import (
    "log"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Person struct {
    ID int `json:"id,omitempty"`
    Firstname string `json:"firstname,omitempty"`
    Lastname string `json:"lastname,omitempty"`
}

type Task struct {
    ID int `json:"id,omitempty"`
    Taskname string `json:"taskname,omitempty"`
}

var tasks []Task

func GetInfo(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(tasks)
}

func UpdateInfo(w http.ResponseWriter, req *http.Request) {

}

func DeleteInfo(w http.ResponseWriter, req *http.Request) {

}

func main() {
    router := mux.NewRouter()
    tasks = append(tasks, Task{ID: 1, Taskname: "Estimation"})
    tasks = append(tasks, Task{ID: 2, Taskname: "Estimate estimation"})
    router.HandleFunc("/users/{username}", GetInfo).Methods("GET")
    router.HandleFunc("/users/{username}", UpdateInfo).Methods("PUT")
    router.HandleFunc("/users/{username}", DeleteInfo).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
