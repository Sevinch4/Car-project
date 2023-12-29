package controller

import (
	"carProject/storage/postgres"
)

type Controller struct {
	Store postgres.Store
}

func New(store postgres.Store) Controller {
	return Controller{
		Store: store,
	}
}
