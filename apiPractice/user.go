// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type SampleDemo struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	Email     string `json:"email"`
// }

// // func GetUser(w http.ResponseWriter, r *http.Request) {

// // 	var a SampleDemo
// // 	a.Firstname = vars["firstname"]
// // 	a.Lastname = vars["lastname"]
// // 	a.email = vars["email"]

// // 	fmt.Fprintf(w, "Firstname: %s \nLastname: %s \nemail: %s \n", a.Firstname, a.Lastname, a.email)

// // }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user SampleDemo
// 	json.NewDecoder(r.Body).Decode(&user)
// 	json.NewEncoder(w).Encode(user)
// }