package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"ocp.com/dino/databaselayer"
)

func main() {

	data := &databaselayer.Animal{
		AnimalType: "Carnotosaurus",
		Nickname:   "carno",
		Zone:       3,
		Age:        13,
	}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)

	resp, err := http.Post("http://localhost:8080/api/dinos/add", "application/json", &b)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(err)
	}

	/*
		resp, err := http.Post("http://localhost:8080/api/dinos/edit/carno", "application/json", &b)
		if err != nil || resp.StatusCode != 200 {
			log.Fatal(err)
		}
	*/
}
