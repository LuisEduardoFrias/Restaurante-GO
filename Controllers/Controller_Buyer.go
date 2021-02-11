package Controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Get a buyer indicating a day of the month
func Get_BuyerPerDay(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "Day")
	fmt.Fprintf(w, "Buyer per day : "+p)
}

//Get a buyer indicating a id
func Get_BuyerPerId(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "Id")
	fmt.Fprintf(w, "Buyer per id : "+p)
}

//Get a list of buyers
func Get_AllBuyers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "listas de compradores")
}
