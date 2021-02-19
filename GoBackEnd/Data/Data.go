package Data

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

func ConnectionToDatabase() *dgo.Txn {

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	op := &api.Operation{
		Schema: `name: string @index(exact) .`,
	}

	ctx := context.Background()

	ctx, _ = context.WithCancel(ctx)

	err1 := dgraphClient.Alter(ctx, op)

	if err1 != nil {
		log.Fatal(err)
	}

	return dgraphClient.NewTxn()
	//defer txn.Discard(ctx)
}
