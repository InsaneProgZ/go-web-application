package repository

import (
	"context"
	"fmt"
	"web-app/src/configuration"
	"web-app/src/model"
)

func SaveCheckout(checkout model.Checkout) {

	client := configuration.MongoConnection()

	result, err :=client.InsertOne(context.TODO(), checkout)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)

}
