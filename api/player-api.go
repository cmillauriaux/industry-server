package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"git.icysoft.fr/cedric/industry-go-server/api/schema"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/xeipuuv/gojsonschema"

	"git.icysoft.fr/cedric/industry-go-server/model"
)

var (
	newPlayerSchema *gojsonschema.Schema
)

func init() {
	var err error
	newPlayerSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(schema.PlayerCreateSchema))
	if err != nil {
		log.Fatal(err)
	}
}

func NewPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Validate datas
	valid, body, err := validateData(newPlayerSchema, r.Body)
	if !valid || err != nil {
		fmt.Fprint(w, structToJSON(makeError(err)))
		return
	}

	// Decode article from body request
	decoder := json.NewDecoder(body)
	var player model.Player
	err = decoder.Decode(&player)
	if err != nil {
		fmt.Fprint(w, structToJSON(makeError(err)))
		return
	}
	defer body.Close()

	// Make a new UID
	player.PlayerID = uuid.NewV4().String()

	// Register player
	err = playerService.NewPlayer(&player)
	if err != nil {
		if err != nil {
			fmt.Fprint(w, structToJSON(makeError(err)))
			return
		}
	}

	w.Write([]byte(structToJSON(Response{Data: player})))
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// get datas
	playerId := mux.Vars(r)["id"]

	// Register player
	player, err := playerService.GetPlayer(playerId)
	if err != nil {
		if err != nil {
			fmt.Fprint(w, structToJSON(makeError(err)))
			return
		}
	}

	w.Write([]byte(structToJSON(Response{Data: player})))
}
