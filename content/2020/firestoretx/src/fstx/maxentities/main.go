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

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		for i := 0; i < 500; i++ {
			key := datastore.NameKey("Sample", fmt.Sprintf("sample%03d", i+1), nil)

			entity := fstx.SampleModel{
				Value:     i + 1,
				UpdatedAt: time.Now(),
			}

			if _, err := tx.Put(key, &entity); err != nil {
				return err
			}
		}

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
