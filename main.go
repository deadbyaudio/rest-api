package main

import (
	"log"
	"net/http"

	"github.com/deadbyaudio/rest-api/controllers"
	"github.com/deadbyaudio/rest-api/models"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", controllers.IndexPeople).Methods("GET")
	router.HandleFunc("/people/{id}", controllers.ShowPerson).Methods("GET")
	router.HandleFunc("/people", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", controllers.DeletePerson).Methods("DELETE")

	models.AddPerson(models.Person{ID: 1, Firstname: "John", Lastname: "Doe", Address: &models.Address{Street: "Carrer de Xátiva", City: "Valencia", Country: "Spain"}})
	models.AddPerson(models.Person{ID: 2, Firstname: "Koko", Lastname: "Doe", Address: &models.Address{City: "Detroit", Country: "USA"}})
	models.AddPerson(models.Person{ID: 3, Firstname: "Francis", Lastname: "Sunday"})

	log.Fatal(http.ListenAndServe(":8080", router))
}
