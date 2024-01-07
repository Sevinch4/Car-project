package main

import (
	"carProject/config"
	"carProject/controller"
	"carProject/storage/postgres"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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

	http.HandleFunc("/driver", con.Driver)
	http.HandleFunc("/car", con.Car)

	fmt.Println("server runnig....")
	http.ListenAndServe("localhost:8080", nil)

	//	var cmdD int
	//	fmt.Print(`
	//			1 - create driver
	//			2 - get driver by id
	//			3 - get list
	//			4 - update driver
	//			5 - delete driver
	//`)
	//	fmt.Scan(&cmdD)
	//
	//	switch cmdD {
	//	case 1:
	//		con.CreateDriver()
	//	case 2:
	//		con.GetDriverByID()
	//	case 3: //getlist
	//		con.GetDriverList()
	//	case 4: //update
	//		con.UpdateDriver()
	//	case 5: //delete
	//		con.DeleteDriver()
	//
	//	}
	//
	//	var cmd int
	//
	//	fmt.Print(`
	//			1 - create car
	//			2 - get car by id
	//			3 - get list
	//			4 - update car
	//			5 - delete car
	//`)
	//	fmt.Scan(&cmd)
	//
	//	switch cmd {
	//	case 1:
	//		con.CreateCar()
	//	case 2:
	//		con.GetCarByID()
	//	case 3:
	//		con.GetCarList()
	//	case 4:
	//		con.UpdateCar()
	//	case 5:
	//		con.DeleteCar()
	//	}

}
