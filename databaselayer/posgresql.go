package databaselayer

import (
	"database/sql"

	_ "github.com/lib/pq" //postgres drive
)

/*PQHandler does ...*/
type PQHandler struct {
	*SQLHandler
}

/*NewPQHandler does ...*/
func NewPQHandler(connection string) (*PQHandler, error) {
	db, err := sql.Open("postgres", connection)
	return &PQHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
