package main

import (
	"context"
	"errors"
	"fmt"
	"fstx"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

var client = fstx.Client()

func mainRollback() {

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		key1 := datastore.NameKey("Sample", "sample1", nil)

		entity1 := fstx.SampleModel{
			Value:     1,
			UpdatedAt: time.Now(),
		}

		if _, err := tx.Put(key1, &entity1); err != nil {
			return err
		}

		key2 := datastore.NameKey("Sample", "sample2", nil)

		entity2 := fstx.SampleModel{
			Value:     2,
			UpdatedAt: time.Now(),
		}

		if _, err := tx.Put(key2, &entity2); err != nil {
			return err
		}

		return errors.New("わざとエラー")
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
