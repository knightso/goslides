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

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		for _, keyName := range []string{"A", "B", "C"} {
			log.Printf("sleeping... %s", sleep)
			time.Sleep(10 * time.Second)

			key := datastore.NameKey("Sample", keyName, nil)
			var entity fstx.SampleModel
			if err := tx.Get(key, &entity); err != nil {
				return err
			}

			log.Printf("key:%s, value:%+v\n", key, entity)
		}

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
