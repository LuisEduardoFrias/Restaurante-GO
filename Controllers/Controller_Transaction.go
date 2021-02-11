package Controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Get a transaction indicating a day of the month
func Get_TransactionPerDay(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "Day")
	fmt.Fprintf(w, "Transaction per day :"+p)
}
