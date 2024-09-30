package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"vector-mongodb-sink/adapters"
)

func StoreJson(b []byte, collection string) {
	var data []map[string]interface{}
	var dataToInsert []interface{}
	err := json.Unmarshal(b, &data)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, m := range data {
		dataToInsert = append(dataToInsert, m)
	}

	fmt.Println(data)
	coll := adapters.GetCollection(adapters.DB, collection)

	if len(dataToInsert) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		_, err := coll.InsertMany(ctx, dataToInsert)

		if err != nil {
			log.Fatal(err)
		}
	}
}
