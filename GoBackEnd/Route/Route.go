package Route

import (
	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Controllers"
	"github.com/go-chi/chi"

	"github.com/go-chi/cors"
)

//Routes
func Routes() *chi.Mux {

	r := chi.NewRouter()

	// CORS básico
	// para obtener más ideas, consulte: https://developer.github.com/v3/#cross-origin-resource-sharing

	r.Use(cors.Handler(cors.Options{

		//AllowOrigins:  	  [] string {"https://foo.com"}, // Use esto para permitir hosts de origen específicos
		AllowedOrigins: []string{"*"},

		//AllowOriginFunc:  func (r * http.Request, cadena de origen) bool {return true},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Valor máximo que ninguno de los principales navegadores ignora
	}))

	//Controller Buyer
	r.Get("/GetAllBuyers", Controllers.Get_AllBuyers)           //Get all buyers
	r.Get("/GetBuyerPerId/{Id}", Controllers.Get_BuyerPerId)    //Get buyer per id
	r.Get("/GetBuyerPerDay/{Day}", Controllers.Get_BuyerPerDay) //Get buyer per day
	r.Post()

	//Controller Product
	r.Get("/GetProductPerDay/{Day}", Controllers.Get_ProductPerDay) //Get product per day

	//Cotroller Transaction
	r.Get("/GetTransactionPerDay/{Day}", Controllers.Get_TransactionPerDay) //Get transaction per day

	return r
}
