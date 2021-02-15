package Route

import (
	"github.com/LuisEduardoFrias/Restaurante/Controllers"
	"github.com/go-chi/chi"
)

//Routes
func Routes() *chi.Mux {

	r := chi.NewRouter()

	//Controller Buyer
	r.Get("/GetAllBuyers", Controllers.Get_AllBuyers)           //Get all buyers
	r.Get("/GetBuyerPerId/{Id}", Controllers.Get_BuyerPerId)    //Get buyer per id
	r.Get("/GetBuyerPerDay/{Day}", Controllers.Get_BuyerPerDay) //Get buyer per day

	//Controller Product
	r.Get("/GetProductPerDay/{Day}", Controllers.Get_ProductPerDay) //Get product per day

	//Cotroller Transaction
	r.Get("/GetTransactionPerDay/{Day}", Controllers.Get_TransactionPerDay) //Get transaction per day

	return r
}
