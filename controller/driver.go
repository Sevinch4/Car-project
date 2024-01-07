package controller

import (
	"carProject/models"
	"carProject/pkg/check"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Controller) Driver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateDriver(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetDriverList(w, r)
		} else {
			c.GetDriverByID(w, r)
		}
	case http.MethodPut:
		c.UpdateDriver(w, r)
	case http.MethodDelete:
		c.DeleteDriver(w, r)
	default:
		fmt.Println("method not allowed")
	}
}

func (c Controller) CreateDriver(w http.ResponseWriter, r *http.Request) {
	driver := models.Driver{}

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		fmt.Println("error is while decoding data", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !check.PhoneNumber(driver.Phone) {
		fmt.Println("phone number format is not correct", driver.Phone)
		handleResponse(w, http.StatusBadRequest, driver)
		return
	}

	id, err := c.Store.DriverStorage.Insert(driver)
	if err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return
	}

	if err != nil {
		fmt.Println("error is while parsing id", err.Error())
		return
	}
	resp, err := c.Store.DriverStorage.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}
	handleResponse(w, http.StatusCreated, resp)
}

func (c Controller) GetDriverByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	driver, err := c.Store.DriverStorage.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, driver)
}

func (c Controller) GetDriverList(w http.ResponseWriter, r *http.Request) {
	drivers, err := c.Store.DriverStorage.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, drivers)
}

func (c Controller) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	driver := models.Driver{}

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		fmt.Println("error is while decoding data", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !check.PhoneNumber(driver.Phone) {
		fmt.Println("phone number format is not correct", driver.Phone)
		handleResponse(w, http.StatusBadRequest, driver)
		return
	}

	if err := c.Store.DriverStorage.Update(driver); err != nil {
		fmt.Println("error is while updating ", err.Error())
		return
	}
	fmt.Println("driver updated!")
	handleResponse(w, http.StatusOK, driver)
}

func (c Controller) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	if err := c.Store.DriverStorage.Delete(id); err != nil {
		fmt.Println("error is while deleting", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)
}
