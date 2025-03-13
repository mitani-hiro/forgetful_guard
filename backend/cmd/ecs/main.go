package main

import (
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

	r := router.NewRouter()

	r.Use(interceptor.Recovery())

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
