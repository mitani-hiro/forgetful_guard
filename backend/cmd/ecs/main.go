package main

import (
	"forgetful-guard/common/caws"
	"forgetful-guard/common/interceptor"
	"forgetful-guard/common/rdb"
	"forgetful-guard/router"
	"log"
)

func main() {

	if err := rdb.InitDB(); err != nil {
		log.Fatal(err)
		return
	}

	if err := caws.NewDynamoDBClient(); err != nil {
		log.Fatalf("failed to initialize DynamoDB client: %v", err)
		return
	}

	r := router.NewRouter()

	r.Use(interceptor.Recovery())

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
