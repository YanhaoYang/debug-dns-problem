package main

import (
	"fmt"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	_, err := elastic.NewClient(
		elastic.SetURL(os.Getenv("ES_SERVER_URL")),
		elastic.SetMaxRetries(15),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to es")
}
