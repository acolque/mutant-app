package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mutant-app/services"
)

func Mutant(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "no se puede leer el request")
		return
	}

	var dna []string
	err2 := json.Unmarshal(reqBody, &dna)
	if err2 != nil {
		respondWithJSON(w, http.StatusBadRequest, "parametro invalido")
		return
	}

	v := new(services.DnaValidator)
	result := v.ValidateDna(dna)
	if !result {
		respondWithJSON(w, http.StatusBadRequest, "dna invalido")
		return
	}

	respondWithJSON(w, http.StatusOK, "app en construccion")
}
