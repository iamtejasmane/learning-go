package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dogs"`
}

var myJson string = `
[
	{
		"first_name": "John",
		"last_name": "Wick",
		"hair_color": "black",
		"has_dogs": true
	},
	{
		"first_name": "John",
		"last_name": "Cena",
		"hair_color": "Gray",
		"has_dogs": false
	}
]`

func main() {
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(unmarshalled)

	var mySlice []Person

	var m1 Person

	m1.FirstName = "John"
	m1.LastName = "Wich"
	m1.HairColor = "black"
	m1.HasDog = true

	var m2 Person
	m2.FirstName = "Carl"
	m2.LastName = "Marx"
	m2.HairColor = "white"
	m2.HasDog = false

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(newJson))
}
