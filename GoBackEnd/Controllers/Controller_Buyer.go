package Controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Models"
	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Models"
	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Repository"
	"github.com/go-chi/chi"
)

//var buyers = []Models.Buyer{
//	{
//		Id:   95195187,
//		Name: "Casanova",
//		Age:  55,
//	},
//	{
//		Id:   90801232,
//		Name: "Skutchan",
//		Age:  76,
//	},
//	{
//		Id:   83516845,
//		Name: "Fakieh",
//		Age:  73,
//	},
//	{
//		Id:   41374564,
//		Name: "Phyllis",
//		Age:  16,
//	},
//	{
//		Id:   34833298,
//		Name: "Gaidano",
//		Age:  66,
//	},
//}

//Get a buyer indicating a day of the month
func Get_BuyerPerDay(w http.ResponseWriter, r *http.Request) {

	d, err := strconv.Atoi(chi.URLParam(r, "Day"))

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(buyers[d])

}

//Get a buyer indicating a id
func Get_BuyerPerId(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "Id"))

	if err != nil {
		log.Fatal(err)
	} else {

		buyer := Models.New_Buyer()

		for i, b := range buyers {
			if b.Id == id {
				buyer = &b
				_ = i
			}
		}

		json.NewEncoder(w).Encode(buyer)

	}
}

//Get a list of buyers
func Get_AllBuyers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "listas de compradores")
}

func Set_Buyers(w http.ResponseWriter, r *http.Request) {

	//Models.Buyer := r.Context().Value("Buyer").(*Models.Buyer)

	Repository.SetBuyer(Models.Buyer{
		Id:   1,
		Name: "joselo",
		Age:  34,
	})
	
}
