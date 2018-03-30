package models

// Person defines a person
type Person struct {
	ID        int      `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address contains the address of a persom
type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

// Person is a crappy data structure given that there is no DB ATM
var people []Person
var idx int

// GetPeople gets all the people in the slice
func GetPeople() []Person {
	return people
}

// AddPerson adds a person to the slice
func AddPerson(p Person) Person {
	idx++
	p.ID = idx
	people = append(people, p)
	return p
}

// FindPersonByID finds a person by id in the slice
func FindPersonByID(id int) Person {
	for _, item := range people {
		if item.ID == id {
			return item
		}
	}
	return Person{}
}

// DestroyPersonByID deletes a person by id from the slice
func DestroyPersonByID(id int) Person {
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			return item
		}
	}
	return Person{}
}
