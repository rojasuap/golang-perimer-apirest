package main

//IMPORTAR PAQUETES
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//ESTRUCTURAS
type Person struct {
	Id        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//FUNCIONES
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range people {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func createPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

	people = append(people, Person{Id: params["id"], FirstName: params["firstname"], LastName: params["lastname"], Address: &Address{City: params["id"], State: "California"}})
}
func deletePeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.Id == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

//FUNCION PRINCIPAL DONDE SE EJECUTA EL CÓDIGO
func main() {
	router := mux.NewRouter()

	people = append(people, Person{Id: "1", FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dubling", State: "California"}})
	people = append(people, Person{Id: "2", FirstName: "Pedro", LastName: "Rojas"})

	//endpoints
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", deletePeopleEndpoint).Methods("DELETE")

	//levantar servidor
	log.Fatal(http.ListenAndServe(":3000", router))
}
