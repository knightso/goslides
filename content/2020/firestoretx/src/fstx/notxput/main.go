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

	log.Println("start")

	// START OMIT
	key := datastore.NameKey("Sample", "test", nil)

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var entity fstx.SampleModel

		if err := tx.Get(key, &entity); err != nil {
			return err
		}

		entity.Value++
		entity.UpdatedAt = time.Now()

		if _, err := client.Put(ctx, key, &entity); err != nil { // HL
			return err
		}

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}
	// END OMIT

	fmt.Printf("done\n")
}
