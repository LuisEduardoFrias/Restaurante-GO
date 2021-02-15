package Controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Get a product indicating a day of the month
func Get_ProductPerDay(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "Day")
	fmt.Fprintf(w, "product per day : "+p)
}
