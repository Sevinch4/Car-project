package controller

import (
	"carProject/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (c Controller) CreateCar() {
	car := getCarInfo()

	if !checkCarInfo(car) {
		return
	}
	id, err := c.Store.CarStorage.Insert(car)
	if err != nil {
		fmt.Println("error is while creating data inside controller err: ", err.Error())
		return
	}
	fmt.Println("id: ", id)
}
func checkCarInfo(car models.Car) bool {
	if car.Year <= 0 || car.Year > time.Now().Year()+1 {
		fmt.Println("year input is not correct")
		return false
	}
	return true
}
func (c Controller) GetCarByID() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error is while parsing id ", err.Error())
		return
	}
	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err.Error())
		return
	}
	fmt.Println("car is : ", car)
}

func (c Controller) GetCarList() {
	cars, err := c.Store.CarStorage.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err)
		return
	}

	fmt.Println(cars)
}

func (c Controller) UpdateCar() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	car := getCarInfo()
	if err := c.Store.CarStorage.Update(car); err != nil {
		fmt.Println("error is while updating", err.Error())
		return
	}

	fmt.Println("data updated")
}

func (c Controller) DeleteCar() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error is while parsing id", err.Error())
		return
	}

	if err = c.Store.CarStorage.Delete(id); err != nil {
		fmt.Println("error is while deleting data", err.Error())
		return
	}

	fmt.Println("data is deleted")
}

func getCarInfo() models.Car {
	var (
		idStr, model, brand, driverID string
		year, cmd                     int
	)
a:
	fmt.Print(`enter command:
					1 - create car
					2 - update car
	`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("input id: ")
		fmt.Scan(&idStr)

		fmt.Print("input model and brand: ")
		fmt.Scan(&model, &brand)

		fmt.Print("input year: ")
		fmt.Scan(&year)

		fmt.Print("input driverID: ")
		fmt.Scan(&driverID)
	} else if cmd == 1 {
		fmt.Print("input model and brand: ")
		fmt.Scan(&model, &brand)

		fmt.Print("input year: ")
		fmt.Scan(&year)

		fmt.Print("input driverID: ")
		fmt.Scan(&driverID)
	} else {
		fmt.Println("not found")
		goto a
	}

	if idStr != "" {
		return models.Car{
			ID:       uuid.MustParse(idStr),
			Model:    model,
			Brand:    brand,
			Year:     year,
			DriverID: uuid.MustParse(driverID),
		}
	}

	return models.Car{
		Model:    model,
		Brand:    brand,
		Year:     year,
		DriverID: uuid.MustParse(driverID),
	}
}
