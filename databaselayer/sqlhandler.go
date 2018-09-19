package databaselayer

import (
	"database/sql"
	"fmt"
)

/*SQLHandler does ...*/
type SQLHandler struct {
	*sql.DB
}

/*GetAllDinos does ...*/
func (handler *SQLHandler) GetAllDinos() ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query("SELECT * FROM Animals")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}

/*GetDinoByNickName does ...*/
func (handler *SQLHandler) GetDinoByNickName(nickname string) (Animal, error) {
	row := handler.QueryRow("SELECT * FROM Animals WHERE nickname = $1", nickname)
	a := Animal{}
	row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, nil
}

/*GetDinosByType does ...*/
func (handler *SQLHandler) GetDinosByType(dinoType string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query("SELECT * FROM Animals WHERE Animal_type = $1", dinoType)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}

/*AddAnimal does ...*/
func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec("INSERT INTO Animals (Animal_type, nickname, zone, age) VALUES ($1, $2, $3, $4)",
		a.AnimalType, a.Nickname, a.Zone, a.Age)
	return err
}

/*UpdateAnimal does ...*/
func (handler *SQLHandler) UpdateAnimal(a Animal, nname string) error {
	_, err := handler.Exec("UPDATE Animals SET Animal_type = $1, nickname = $2, zone = $3, age = $4 WHERE nickname = $5",
		a.AnimalType, a.Nickname, a.Zone, a.Age, nname)
	return err
}
