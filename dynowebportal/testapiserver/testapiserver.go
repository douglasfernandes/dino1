package main

import (
	"fmt"
	"log"

	"ocp.com/dino/databaselayer"
	"ocp.com/dino/dynowebportal/dinoapi"
)

// Configurações de banco de dados
const (
	DBHOST = "localhost"
	DBPORT = 5432
	DBUSER = "postgres"
	DBPWD  = "123456"
	DBNAME = "dino"
)

func main() {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPWD, DBNAME)

	db, err := databaselayer.GetDatabaseHandler(databaselayer.POSTGRESQL, dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	dinoapi.RunAPI("localhost:8080", db)

}
