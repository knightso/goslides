package main

import (
	"context"
	"flag"
	"fmt"
	"fstx"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

var client = fstx.Client()

func main() {
	ctx := context.Background()

	var value = flag.Int("v", 1, "Value")

	flag.Parse()

	keyNames := flag.Args()

	log.Printf("putting value:%d into entities %v\n", *value, keyNames)

	keys := make([]*datastore.Key, len(keyNames))
	entities := make([]*fstx.SampleModel, len(keyNames))

	for i, keyName := range keyNames {
		keys[i] = datastore.NameKey("Sample", keyName, nil)
		entities[i] = &fstx.SampleModel{
			Value:     *value,
			UpdatedAt: time.Now(),
		}
	}

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {

		if _, err := tx.PutMulti(keys, entities); err != nil {
			return err
		}

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
