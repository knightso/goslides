package fstx

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

var client *datastore.Client

func init() {
	projectID := os.Getenv("STORE_PROJECT")
	fmt.Println(projectID)
	clt, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatal(err)
	}
	client = clt
}

func Client() *datastore.Client {
	return client
}

// SampleModel is sample model.
type SampleModel struct {
	Value     int
	UpdatedAt time.Time
}
