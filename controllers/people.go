package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deadbyaudio/rest-api/models"

	"github.com/gorilla/mux"
)

// IndexPeople Shows all the people in the system
func IndexPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.GetPeople())
}

// ShowPerson Shows the person referenced by the given id
func ShowPerson(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	p := models.FindPersonByID(i)

	if (models.Person{}) == p {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

// CreatePerson Creates a new person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p models.Person
	_ = json.NewDecoder(r.Body).Decode(&p)

	p = models.AddPerson(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// DeletePerson Deletes a person
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	p := models.DestroyPersonByID(i)

	if (models.Person{}) == p {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}
