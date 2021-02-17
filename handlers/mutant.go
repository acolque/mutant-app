package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mutant-app/business"
	"github.com/mutant-app/services"
)

type DnaRequest struct {
	Dna []string
}

func Mutant(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "No se puede leer el request")
		return
	}

	var request DnaRequest
	err2 := json.Unmarshal(reqBody, &request)
	if err2 != nil {
		respondWithJSON(w, http.StatusBadRequest, "Parametro invalido")
		return
	}

	dna := request.Dna
	validator := new(services.DnaValidator)
	result := validator.ValidateDna(dna)
	if !result {
		respondWithJSON(w, http.StatusBadRequest, "DNA invalido")
		return
	}

	detector := services.NewDnaMutantDetector()
	// db := new(services.MutantMockDb)
	db := services.NewMutantMongodb()
	myBusiness := business.NewMutantBusiness(detector, db)
	isMutant, msg := myBusiness.IsMutant(dna)
	if isMutant {
		respondWithJSON(w, http.StatusOK, fmt.Sprintf("Es mutante. %s", msg))
	} else {
		respondWithJSON(w, http.StatusForbidden, fmt.Sprintf("No es mutante. %s", msg))
	}
}
