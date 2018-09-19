package dinoapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"ocp.com/dino/databaselayer"
)

/*DinoRESTAPIHandler does ...*/
type DinoRESTAPIHandler struct {
	dbhandler databaselayer.DinoBDHandler
}

func newDinoRESTAPIHandler(db databaselayer.DinoBDHandler) *DinoRESTAPIHandler {
	return &DinoRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *DinoRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Nenhum criterio de busca encontrado, você pode pesquisar tanto por nickname via /api/dinos/nickname/rex , ou
						pesquisar por type via /api/dinos/type/velociraptor`)
		return
	}

	searchkey, ok := vars["Search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Nenhum criterio de busca encontrado, você pode pesquisar tanto por nickname via /api/dinos/nickname/rex , ou
						pesquisar por type via /api/dinos/type/velociraptor`)
		return
	}

	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error

	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetDinoByNickName(searchkey)
	case "type":
		animals, err = handler.dbhandler.GetDinosByType(searchkey)
	case "all":
		animals, err = handler.dbhandler.GetAllDinos()
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Ocorreu erro durante a pesquisa na base de dados: %v ", err)
		return
	}
	if len(animals) > 0 {
		json.NewEncoder(w).Encode(animals)
		return
	}
	json.NewEncoder(w).Encode(animal)

}

func (handler *DinoRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["Operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Nenhum criterio de busca encontrado, você pode pesquisar tanto por nickname via /api/dinos/nickname/rex , ou
						pesquisar por type via /api/dinos/type/velociraptor`)
		return
	}
	var animal databaselayer.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Não foi possível decodificar o corpo da requisição para o formato json %v", err)
		return
	}
	switch strings.ToLower(operation) {
	case "add":
		err = handler.dbhandler.AddAnimal(animal)
	case "edit":
		nickname := r.RequestURI[len("/api/dinos/edit/"):]
		log.Println("Request... de Edição por Nickname:", nickname)
		err = handler.dbhandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ocorreu um erro durante o processamento da requisição %v", err)
	}
}
