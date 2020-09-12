package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"fstx"
	"log"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

var client = fstx.Client()

// ok
func mainOk() {

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

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}

// rollback
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

// over 25 eg
//func main_over25eg() {
func main25eg() {

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		for i := 0; i < 26; i++ {
			key := datastore.NameKey("Sample", fmt.Sprintf("sample%02d", i+1), nil)

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

// over 500 entities
func mainOver500() {

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		for i := 0; i < 501; i++ {
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

// 500put+1del
func main500put1del() {

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		for i := 0; i < 499; i++ {
			key := datastore.NameKey("Sample", fmt.Sprintf("sample3-%03d", i+1), nil)

			entity := fstx.SampleModel{
				Value:     i + 1,
				UpdatedAt: time.Now(),
			}

			if _, err := tx.Put(key, &entity); err != nil {
				return err
			}
		}

		key := datastore.NameKey("Sample", fmt.Sprintf("sample%03d", 1), nil)
		if err := tx.Delete(key); err != nil {
			return err
		}

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}

// 1EGで501
func main501In1Eg() {

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		pkey := datastore.IDKey("Parent", 1, nil)
		for i := 0; i < 501; i++ {
			key := datastore.NameKey("Sample", fmt.Sprintf("sample%03d", i+1), pkey)

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

// non-transactional put
func mainNoTxPut(keyName string, value int) {
	//func main() {

	log.Println("start")

	ctx := context.Background()

	key := datastore.NameKey("Sample", keyName, nil)

	entity := fstx.SampleModel{
		Value:     value,
		UpdatedAt: time.Now(),
	}

	if _, err := client.Put(ctx, key, &entity); err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}

// locks
//func main_locks() {
func main() {

	var sleepBefore = flag.Duration("sb", 0, "sleep duration before operation")
	var sleepAfter = flag.Duration("sa", 0, "sleep duration after operation")
	var keyName = flag.String("k", "test", "Key name")
	var value = flag.Int("v", 1, "Value")
	var noTx = flag.Bool("n", false, "Non-transactional Put")

	flag.Parse()

	if *noTx {
		// Non-transactional Put
		mainNoTxPut(*keyName, *value)
		return
	}

	operation := flag.Arg(0) // Get, Put, GetPut, Insert, Update, Upsert

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
		default:
			return fmt.Errorf("illegal operation:%s", operation)
		}

		log.Printf("sleeping... %s\n", *sleepAfter)
		time.Sleep(*sleepAfter)

		log.Println("commiting...")

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	log.Printf("done.")
}

// lock get&put
func mainLockGetput() {

	var sleep = flag.Duration("s", 0, "sleep duration")

	flag.Parse()

	fmt.Println(sleep)

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		key1 := datastore.NameKey("Sample", "sample1", nil)

		var entity fstx.SampleModel

		if err := tx.Get(key1, &entity); err != nil {
			return err
		}

		fmt.Printf("updating %+v\n", entity)

		entity.Value++
		entity.UpdatedAt = time.Now()

		time.Sleep(*sleep)

		if _, err := tx.Put(key1, &entity); err != nil {
			return err
		}

		fmt.Printf("updated %+v\n", entity)

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}

// lock only get vs get&put
func mainGetVsGetPut() {

	var sleep = flag.Duration("s", 0, "sleep duration")
	var put = flag.Bool("p", false, "put entity")

	flag.Parse()

	fmt.Println(sleep)

	ctx := context.Background()

	if _, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		key1 := datastore.NameKey("Sample", "sample1", nil)

		var entity fstx.SampleModel

		if err := tx.Get(key1, &entity); err != nil {
			return err
		}

		if *put {
			fmt.Printf("updating %+v\n", entity)

			entity.Value++
			entity.UpdatedAt = time.Now()

			if _, err := tx.Put(key1, &entity); err != nil {
				return err
			}

			fmt.Printf("updated %+v\n", entity)
		}

		fmt.Printf("sleeping %s\n", *sleep)
		time.Sleep(*sleep)

		return nil
	}, datastore.MaxAttempts(1)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}

// クエリ＆カーソル
func mainQueryCursor() {

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
