package controller

import (
	"carProject/models"
	"carProject/pkg/check"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Controller) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCar(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			c.GetCarList(w, r)
		} else {
			c.GetCarByID(w, r)
		}
	case http.MethodPut:
		c.UpdateCar(w, r)
	case http.MethodDelete:
		c.DeleteCar(w, r)

	}
}

func (c Controller) CreateCar(w http.ResponseWriter, r *http.Request) {
	car := models.Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		fmt.Println("error is while decoding data", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(car.Year)
	if err := check.Year(car.Year); err != nil {
		fmt.Println("error is while input year", car.Year)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.CarStorage.Insert(car)
	if err != nil {
		fmt.Println("error is while creating data inside controller err: ", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}
func (c Controller) GetCarByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, car)
}

func (c Controller) GetCarList(w http.ResponseWriter, r *http.Request) {
	cars, err := c.Store.CarStorage.GetList()

	if err != nil {
		fmt.Println("error is while getting list", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) UpdateCar(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	car := models.Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		fmt.Println("error is while decoding data")
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := check.Year(car.Year); err != nil {
		fmt.Println("year format is not correct", car.Year)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if car.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, car)
		return
	}

	car.ID = id

	if err := c.Store.CarStorage.Update(car); err != nil {
		fmt.Println("error is while updating", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Println("eror is while getting by id")
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) DeleteCar(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Store.CarStorage.Delete(id); err != nil {
		fmt.Println("error is while deleting data", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)
	fmt.Println("data is deleted")
}
