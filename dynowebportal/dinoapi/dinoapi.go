package dinoapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"ocp.com/dino/databaselayer"
)

/*RunAPI does ...
HTTP GET -> Para buscar por -> api/dinos/nickname/rex ou api/dinos/type/velociraptor
HTTP POST -> Para adicionar -> api/dinos/add/ ou para editar api/dinos/edit/
*/
func RunAPI(endpoint string, db databaselayer.DinoBDHandler) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

/*RunAPIOnRouter does ...*/
func RunAPIOnRouter(r *mux.Router, db databaselayer.DinoBDHandler) {
	handler := newDinoRESTAPIHandler(db)
	apirouter := r.PathPrefix("/api/dinos").Subrouter()
	apirouter.Methods("GET").Path("/{SearchCriteria}/{Search}").HandlerFunc(handler.searchHandler)
	apirouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)
}
