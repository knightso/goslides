package main

import (
	"context"
	"fmt"
	"fstx"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

var client = fstx.Client()

func main() {
	ctx := context.Background()

	sleep := 10 * time.Second

	keyA := datastore.NameKey("Sample", "A", nil)
	keyC := datastore.NameKey("Sample", "C", nil)

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		for i := 0; i < 3; i++ {
			q := datastore.NewQuery("Sample").
				Filter("__key__ >=", keyA).
				Filter("__key__ <=", keyC).
				Transaction(tx)

			var entities []*fstx.SampleModel
			if _, err := client.GetAll(ctx, q, &entities); err != nil {
				return err
			}

			for j, entity := range entities {
				log.Printf("%d: %+v\n", j, entity)
			}

			log.Printf("sleeping... %s", sleep)
			time.Sleep(10 * time.Second)
		}

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
