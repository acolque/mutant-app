package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mutant-app/business"
	"github.com/mutant-app/services"
)

func Mutant(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "No se puede leer el request")
		return
	}

	var dna []string
	err2 := json.Unmarshal(reqBody, &dna)
	if err2 != nil {
		respondWithJSON(w, http.StatusBadRequest, "Parametro invalido")
		return
	}

	v := new(services.DnaValidator)
	result := v.ValidateDna(dna)
	if !result {
		respondWithJSON(w, http.StatusBadRequest, "DNA invalido")
		return
	}

	detector := new(services.DnaMutantDetector)
	myBusiness := business.NewMutantBusiness(detector)
	isMutant := myBusiness.IsMutant(dna)
	if isMutant {
		respondWithJSON(w, http.StatusOK, "Es mutante")
	} else {
		respondWithJSON(w, http.StatusForbidden, "No es mutante")
	}
}
