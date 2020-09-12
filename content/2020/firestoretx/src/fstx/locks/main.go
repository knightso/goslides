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

// locks
func main() {

	var sleepBefore = flag.Duration("sb", 0, "sleep duration before operation")
	var sleepAfter = flag.Duration("sa", 0, "sleep duration after operation")
	var keyName = flag.String("k", "test", "Key name")
	var value = flag.Int("v", 1, "Value")

	flag.Parse()

	operation := flag.Arg(0) // Get, Put, GetPut, Insert, Update, Upsert, GetUpdate, Query

	ctx := context.Background()

	log.Println("start")

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {

		log.Printf("sleeping... %s\n", *sleepBefore)
		time.Sleep(*sleepBefore)

		key := datastore.NameKey("Sample", *keyName, nil)

		switch operation {
		case "Get":
			log.Println("getting...")

			var entity fstx.SampleModel
			if err := tx.Get(key, &entity); err != nil {
				return err
			}

			log.Printf("got %+v\n", entity)
		case "Put":
			log.Println("putting...")

			entity := fstx.SampleModel{
				Value:     *value,
				UpdatedAt: time.Now(),
			}
			if _, err := tx.Put(key, &entity); err != nil {
				return err
			}
		case "GetPut":
			log.Println("getting...")

			var entity fstx.SampleModel
			if err := tx.Get(key, &entity); err != nil && err != datastore.ErrNoSuchEntity {
				return err
			}

			log.Printf("got %+v\n", entity)

			log.Println("putting...")

			entity.Value = *value
			entity.UpdatedAt = time.Now()

			if _, err := tx.Put(key, &entity); err != nil {
				return err
			}
		case "Insert":
			log.Println("inserting...")

			entity := fstx.SampleModel{
				Value:     *value,
				UpdatedAt: time.Now(),
			}
			if _, err := tx.Mutate(datastore.NewInsert(key, &entity)); err != nil {
				return err
			}
		case "Update":
			log.Println("updating...")

			entity := fstx.SampleModel{
				Value:     *value,
				UpdatedAt: time.Now(),
			}
			if _, err := tx.Mutate(datastore.NewUpdate(key, &entity)); err != nil {
				return err
			}
		case "Upsert":
			log.Println("upserting...")

			entity := fstx.SampleModel{
				Value:     *value,
				UpdatedAt: time.Now(),
			}
			if _, err := tx.Mutate(datastore.NewUpsert(key, &entity)); err != nil {
				return err
			}
		case "GetUpdate":
			var entity fstx.SampleModel
			if err := tx.Get(key, &entity); err != nil && err != datastore.ErrNoSuchEntity {
				return err
			}

			log.Printf("got %+v\n", entity)

			entity.Value = *value
			entity.UpdatedAt = time.Now()

			log.Println("updating...")

			if _, err := tx.Mutate(datastore.NewUpdate(key, &entity)); err != nil {
				return err
			}
		case "Query":
			keyA := datastore.NameKey("Sample", "A", nil)
			keyC := datastore.NameKey("Sample", "C", nil)

			q := datastore.NewQuery("Sample").
				Filter("__key__ >=", keyA).
				Filter("__key__ <=", keyC).
				Transaction(tx)

			log.Println("querying...")

			var entities []*fstx.SampleModel
			if _, err := client.GetAll(ctx, q, &entities); err != nil {
				return err
			}

			for j, entity := range entities {
				log.Printf("%d: %+v\n", j, entity)
			}
		default:
			return fmt.Errorf("illegal operation:%s", operation)
		}

		log.Printf("sleeping... %s\n", *sleepAfter)
		time.Sleep(*sleepAfter)

		log.Println("committing...")

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	log.Printf("done.")
}
