package api

import (
	"encoding/json"
	"fmt"
	"github.com/DavidSkeppstedt/Automa/db"
	"github.com/DavidSkeppstedt/Automa/model"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"os/exec"
	"log"
)

var (
	actions = map[string]struct{}{"on": {}, "off": {}}
)

func apiHandler(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	fmt.Fprintf(response, "Automa API. Version 1.0")
}

func lampHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//fetch array of lamps from db.
	lamp, _ := strconv.Atoi(ps.ByName("lamp"))
	exist, err := db.LampExists(lamp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Database error")
		return
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Lamp not found")
		return
	}
	aLamp, _ := db.GetLamp(lamp)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aLamp)
}

func createLampHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not handle request body")
		return
	}
	var newLamp model.Lamp
	if err = json.Unmarshal(body, &newLamp); err != nil {
		w.WriteHeader(422) // unprocessable entire.
		json.NewEncoder(w).Encode("Could not parse json")
		return
	}

	db.AddLamp(newLamp)
	w.WriteHeader(http.StatusCreated)

}

func lampActionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//check if the lamp exist
	lamp, _ := strconv.Atoi(ps.ByName("lamp"))
	exist, err := db.LampExists(lamp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Database error")
		return
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Lamp not found")
		return
	}

	//check if action is okay.
	action := ps.ByName("action")
	_, ok := actions[action]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Action not accepted. Use on or off.")
		return
	}

	//here we should execute the command to the 433 Mhz controller.
	aLamp,_ := db.GetLamp(lamp)
	switch action {
	case "on":
		log.Println("Hej")
		//~/dev/lamp/
		exec.Command("/bin/sh", "-c", "~/dev/lamp/./send "+strconv.Itoa(aLamp.Id) +" 1").Output()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Turned on")
	case "off":
		log.Println("hej då")
		out,_ := exec.Command("/bin/sh", "-c", "~/dev/lamp/./send "+strconv.Itoa(aLamp.Id)+" 0").Output()
		log.Println("Lamp said:",string(out))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Turned off")
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func lampsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//fetch array of lamps from db.
	lamps, err := db.FetchLamps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Database error")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lamps)
}
