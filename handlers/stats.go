package handlers

import (
	"net/http"

	"github.com/mutant-app/business"
	"github.com/mutant-app/services"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	db := services.NewMutantMongodb()
	myBusiness := business.NewStatsBusiness(db)

	result, err := myBusiness.GetStats()
	if err != nil {
		respondWithJSON(w, http.StatusServiceUnavailable, "Hubo un error al obtener las estadisticas")
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}
