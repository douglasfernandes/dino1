package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"ocp.com/dino/databaselayer"
	"ocp.com/dino/dynowebportal"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
	DbHost        string `json:"dbhost"`
	DbPort        string `json:"dbport"`
	DbUser        string `json:"dbuser"`
	DbPwd         string `json:"dbpwd"`
	DbName        string `json:"dbname"`
	FrontEnd      string `json:"frontend"`
}

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := new(configuration)
	json.NewDecoder(file).Decode(config)

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPwd, config.DbName)

	log.Println("Iniciando o servidor web no endereço ", config.ServerAddress)
	log.Println(dynowebportal.RunWebPortal(databaselayer.POSTGRESQL, config.ServerAddress, dbinfo, config.FrontEnd))
	//dynowebportal.RunWebPortal(databaselayer.POSTGRESQL, config.ServerAddress, dbinfo, config.FrontEnd)
}
