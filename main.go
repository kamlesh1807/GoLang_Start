package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}
type Users []user

func alluser(w http.ResponseWriter, r *http.Request) {
	users := Users{
		user{Name: "kamlesh", Email: "Kamlesh@gla.ac.in"},
		user{Name: "yadav", Email: "yadav@gmail.com"},
	}
	fmt.Println("ALlnEndpointhit : all user")
	json.NewEncoder(w).Encode(users)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", alluser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
