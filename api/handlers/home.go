package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mutant-app/api/services"
)

type UrlDel struct {
	Url      string
	ShortUrl string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := new(services.Person)
	nombre := p.GetName()

	// welcome := "Hola mundo! yo soy " + nombre + " Y esta es mi 1ra app en GO"

	welcome := struct {
		MiNombre   string
		Mensaje    string
		VersionApp string
	}{
		"Hola mundo! yo soy " + nombre,
		"Y esta es mi 1ra app en GO con DOCKER",
		"v1.0",
	}

	respondWithJSON(w, http.StatusOK, welcome)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
