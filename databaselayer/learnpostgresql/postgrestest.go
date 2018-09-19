package main

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
)

// Configurações de banco de dados
const (
	DBHOST = "localhost"
	DBPORT = 5432
	DBUSER = "postgres"
	DBPWD  = "123456"
	DBNAME = "dino"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPWD, DBNAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from animals")
	handlerows(rows, err)

	testTransaction(db)

}

func handlerows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
}

func testTransaction(db *sql.DB) {
	fmt.Println("........Transaction Begin.............")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(9)
	handlerows(rows, err)

	rows, err = stmt.Query(15)
	handlerows(rows, err)

	results, err := tx.Exec("UPDATE animals SET age = $1 WHERE id = $2", 17, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.RowsAffected())

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("........Transaction End.............")
}
