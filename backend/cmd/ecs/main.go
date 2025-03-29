package main

import (
	"forgetful-guard/common/caws"
	"forgetful-guard/common/interceptor"
	"forgetful-guard/common/logger"
	"forgetful-guard/common/rdb"
	"forgetful-guard/router"
)

func init() {
	logger.InitLogger()
}

func main() {

	if err := rdb.InitDB(); err != nil {
		logger.Error("DB initialization error", err)
		return
	}

	if err := caws.NewDynamoDBClient(); err != nil {
		logger.Error("DynamoDB client initialization error", err)
		return
	}

	if err := caws.NewLocationClient(); err != nil {
		logger.Error("Location Service client initialization error", err)
		return
	}

	r := router.NewRouter()

	r.Use(interceptor.Recovery())

	if err := r.Run(":8080"); err != nil {
		logger.Error("gin run error", err)
		return
	}
}
