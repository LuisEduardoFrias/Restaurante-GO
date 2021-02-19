package Repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/LuisEduardoFrias/Restaurante/Models"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

func SetBuyer(d *dgo.Txn, b *Models.Buyer) {

	ctx := context.Background()

	ctx, _ = context.WithCancel(ctx)

	pb, err := json.Marshal(b)

	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson: pb,
	}

	res, err := d.Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
