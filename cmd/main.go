package main

import (
	"carProject/config"
	"carProject/controller"
	"carProject/storage/postgres"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}

	defer store.DB.Close()

	con := controller.New(store)

	//con.CreateCar()
	//con.GetCarByID()
	con.GetCarList()
	//con.UpdateCar()
	con.DeleteCar()
}
