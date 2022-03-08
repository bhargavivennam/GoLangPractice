// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage hi!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }

// func main() {
// 	handleRequests()
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
