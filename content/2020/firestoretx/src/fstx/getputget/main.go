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

	key := datastore.NameKey("Sample", "test", nil)

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var entity fstx.SampleModel

		if err := tx.Get(key, &entity); err != nil {
			return err
		}

		log.Printf("got before put: %+v", entity)

		entity.Value++
		entity.UpdatedAt = time.Now()

		if _, err := tx.Put(key, &entity); err != nil {
			return err
		}

		entity = fstx.SampleModel{} // clear
		if err := tx.Get(key, &entity); err != nil {
			return err
		}

		log.Printf("got after put: %+v", entity)

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
