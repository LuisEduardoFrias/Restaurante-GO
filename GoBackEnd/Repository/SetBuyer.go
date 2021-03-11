//package Repository
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//
//	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Data"
//	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Models"
//	"github.com/dgraph-io/dgo/protos/api"
//)
//
//func SetBuyer(b *Models.Buyer) {
//
//	d := Data.ConnectionToDatabase()
//
//	ctx := context.Background()
//
//	ctx, _ = context.WithCancel(ctx)
//
//	//pb, err := json.Marshal(b)
//
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	buyer, err := json.Marshal(b)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	mu := *api.Mutation{
//		SetJson: buyer,
//	}
//
//	res, err := d.Mutate(ctx, mu)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(res)
//
//	fmt.Printf(b.Name)
//}
