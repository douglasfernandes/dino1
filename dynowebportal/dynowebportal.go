package dynowebportal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"ocp.com/dino/databaselayer"
	"ocp.com/dino/dynowebportal/dinoTemplate"
	"ocp.com/dino/dynowebportal/dinoapi"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*DataFeedMessage does ... */
type DataFeedMessage struct {
	Heartrate     int
	Bloodpressure int
}

/*
//RunWebPortal ponto de entrada para nossa aplicacao
func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, nil)
}
*/

//RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(dbtype uint8, addr, dbconnection, frontend string) error {
	rand.Seed(time.Now().UTC().UnixNano())
	r := mux.NewRouter()
	db, err := databaselayer.GetDatabaseHandler(dbtype, dbconnection)
	if err != nil {
		return err
	}
	dinoapi.RunAPIOnRouter(r, db)

	r.Path("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		dinoTemplate.Homepage("Dino Portal", "Welcome to the Dino portal, where you can find metrics and information ...", w)
	})

	r.PathPrefix("/metrics/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		animals, err := db.GetAllDinos()
		if err != nil {
			log.Print(err)
			return
		}
		dinoTemplate.HandleMetrics(animals, w)
	})

	r.PathPrefix("/info/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		animals, err := db.GetAllDinos()
		if err != nil {
			log.Print(err)
			return
		}
		dinoTemplate.HandleInfo(animals, w)
	})

	fileserver := http.FileServer(http.Dir(frontend))
	r.Path("/dinodatafeed").HandlerFunc(dinoDataFeedHandler)

	r.PathPrefix("/").Handler(fileserver)
	return http.ListenAndServe(addr, r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem Vindo ao Park dos Dinossauros %s", r.RemoteAddr)
}

func dinoDataFeedHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Could not establish websocket connection, error", err)
		return
	}
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()

		if err == io.EOF {
			//socket fechado
			break
		}

		if err != nil {
			log.Println("Could not read message from websocket, error", err)
			return
		}

		active := true
		if messageType == websocket.CloseMessage {
			log.Println("closing websocket... ")
			active = false
			break
		}

		go func(dino string) {

			for active {
				time.Sleep(1 * time.Second)
				msg := &DataFeedMessage{rand.Intn(300) + 1, rand.Intn(1000) + 300}
				//msg := dino + strconv.Itoa(rand.Intn(300)+1)
				databytes, err := json.Marshal(msg)
				if err != nil {
					log.Println("Could not convert data to JSON, error", databytes)
					return
				}
				if err = conn.WriteMessage(messageType, databytes); err != nil {
					log.Println("Could not write message to websocket, error", err)
					return
				}
			}
		}(string(p))
	}
}
