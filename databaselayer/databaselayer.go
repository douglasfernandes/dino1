package databaselayer

import (
	"errors"
)

//Tipos de base de dados suportados
const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

/*Animal does ...*/
type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

/*DinoBDHandler does ...*/
type DinoBDHandler interface {
	GetAllDinos() ([]Animal, error)
	GetDinoByNickName(string) (Animal, error)
	GetDinosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

var errDbTypeNotSupported = errors.New("O tipo provido para a base de dados não é suportado")

/*GetDatabaseHandler does ...*/
func GetDatabaseHandler(dbType uint8, connection string) (DinoBDHandler, error) {
	switch dbType {
	case POSTGRESQL:
		return NewPQHandler(connection)
	}
	return nil, errDbTypeNotSupported

}
