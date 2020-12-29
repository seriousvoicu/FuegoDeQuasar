package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/seriousvoicu/FuegoDeQuasar/db"
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/satellites"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

var satellitieManager *satellites.Satmanager

type Response struct {
	Pos     vectors.Vector2 `json:"position"`
	Message string          `json:"message"`
}

func CreateClusterFromJson(w http.ResponseWriter, r *http.Request) {
	repo := db.SatellitiesRepo{}
	//repo := db.SatellitiesRepoEng{}
	satellitieManager.InstantiateClusterFromJson(r.Body, &repo)

	defer r.Body.Close()

	message := satellitieManager.GetMessage()

	position := satellitieManager.GetLocation()

	w.Header().Set("Content-Type", "application/json")

	if exestate.OnError(satellitieManager) {
		exestate.StateEvalJson(w, satellitieManager)
		return
	}

	response := Response{Pos: position, Message: message}

	json.NewEncoder(w).Encode(response)
}

//Instancia los satelites desde la base, les setea las posiciones especificadas y realiza la triangulacion
func GetLocation(distances ...float32) (x, y float32) {
	repo := db.SatellitiesRepo{}
	//repo := db.SatellitiesRepoEng{}

	satellitieManager.InstantiateClusterFromDB(&repo)

	satellitieManager.SetClusterDistances(distances)

	point := satellitieManager.GetLocation()

	if exestate.OnError(satellitieManager) {
		fmt.Println(satellitieManager.GetState(true).UserError)
		return
	}

	return float32(point.X), float32(point.Y)
}

//Instancia los satelites desde la base, les setea los mensajes
func GetMessage(messages ...[]string) (msg string) {

	repo := db.SatellitiesRepo{}
	//repo := db.SatellitiesRepoEng{}

	satellitieManager.InstantiateClusterFromDB(&repo)

	satellitieManager.SetClusterMessages(messages)

	message := satellitieManager.GetMessage()

	if exestate.OnError(satellitieManager) {
		fmt.Println(satellitieManager.GetState(true).UserError)
		return
	}

	return message
}

func main() {
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/create", CreateClusterFromJson)

	/*a := &api{}
	r := mux.NewRouter()
	r.HandleFunc("/", a.CreateClusterFromJson).Methods(http.MethodPost)

	a.router = r

	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))*/

	satellitieManager = &satellites.Satmanager{}

	http.HandleFunc("/topsecret", CreateClusterFromJson)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
