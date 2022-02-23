package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// // func homePage(w http.ResponseWriter, r *http.Request) {
// // 	// fmt.Fprintf(w, "Endpoint called: homePage()")
// // 	fmt.Fprintf(w, "/books/{title}/page/{page} : homePage()")
// // 	vars := mux.Vars(r)
// // 	title := vars["title"]
// // 	page := vars["page"]

// // 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// // }

// func handleRequests() {
// 	r := mux.NewRouter()

// 	// r.HandleFunc("/", homePage).Methods("GET")
// 	r.HandleFunc("/addition/{a}/{b}", func(w http.ResponseWriter, r *http.Request) {

// 		// vars := mux.Vars(r)
// 		var a, b, c int
// 		a = 2
// 		b = 4

// 		c = a + b
// 		fmt.Fprintf(w, "sum: %d \n", c)

// 		// fmt.Fprintf(w, "You've requested the sum:  %s on page \n", r.c)
// 	}).Methods("GET")
// 	// 	r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {

// 	// 		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", "i", "page")
// 	// 	}).Methods("POST")

// 	r.HandleFunc("/pointers", func(w http.ResponseWriter, r *http.Request) {

// 		type Sample struct {
// 			Id     int
// 			Salary int
// 		}

// 		var s Sample
// 		s.Id = 1
// 		s.Salary = 100000
// 		fmt.Fprintf(w, "ID: %d Salary: %d \n", s.Id, s.Salary)

// 		// s.Id = 3
// 		// fmt.Println(s.Id)

// 	}).Methods("Get")
type SampleDemo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/Sampleapi/{firstname}/{lastname}/{email}/details", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		var a SampleDemo
		a.Firstname = vars["firstname"]
		a.Lastname = vars["lastname"]
		a.Email = vars["email"]
		// a.Firstname = "bhargavi"
		// a.Lastname = "Vr"
		// a.email = "b@gmail.com"

		fmt.Fprintf(w, "Firstname: %s \nLastname: %s \nemail: %s \n", a.Firstname, a.Lastname, a.Email)

	}).Methods("POST")

	// }

	// r.HandleFunc("/users/{Id}", CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", r))

}

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user SampleDemo
// 	json.NewDecoder(r.Body).Decode(&user)
// 	json.NewEncoder(w).Encode(user)
// }

func main() {
	handleRequests()

}
