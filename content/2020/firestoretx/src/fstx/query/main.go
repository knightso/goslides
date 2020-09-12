package main

import (
	"context"
	"fmt"
	"fstx"
	"log"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

var client = fstx.Client()

func main() {
	ctx := context.Background()

	q := datastore.NewQuery("Sample").Limit(100)

	total := 0
	var cursor datastore.Cursor

	for {
		count := 0
		var tmpCur datastore.Cursor

		if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) (err error) {
			q := q.Transaction(tx)

			if cursor.String() != "" {
				q = q.Start(cursor)
			}

			it := client.Run(ctx, q)

			keys := make([]*datastore.Key, 0, 100)
			entities := make([]*fstx.SampleModel, 0, 100)

			for {
				entity := &fstx.SampleModel{}
				key, err := it.Next(entity)
				if err != nil {
					if err == iterator.Done {
						break
					}
					return err
				}

				// update
				entity.UpdatedAt = time.Now()

				keys = append(keys, key)
				entities = append(entities, entity)
			}

			if len(keys) > 0 {
				tx.PutMulti(keys, entities)
			}

			tmpCur, err = it.Cursor()
			if err != nil {
				return err
			}

			count = len(keys)

			return nil
		}); err != nil {
			log.Fatal(err)
		}

		total += count
		cursor = tmpCur

		if count < 100 {
			break
		}
	}

	fmt.Printf("done. total:%d\n", total)
}
