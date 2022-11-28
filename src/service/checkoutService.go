package service

import (
	"web-app/src/model"
	"web-app/src/repository"
)

func SaveCheckout(checkout model.Checkout) {

	repository.SaveCheckout(checkout)
}
