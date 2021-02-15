package Server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Localhost(r *chi.Mux) {

	port := ":3000"

	//localhost
	fmt.Println("Serving on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
