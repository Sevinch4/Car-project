package controller

import (
	"carProject/models"
	"fmt"
	"github.com/google/uuid"
)

func (c Controller) CreateDriver() {
	car := getDriverInfo()

	id, err := c.Store.DriverStorage.Insert(car)
	if err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return
	}
	fmt.Println("id: ", id)
	fmt.Println("driver added!")
}

func (c Controller) GetDriverByID() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error is while parsing id", err.Error())
		return
	}
	driver, err := c.Store.DriverStorage.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}
	fmt.Println("driver is: ", driver)
}

func (c Controller) GetDriverList() {
	drivers, err := c.Store.DriverStorage.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		return
	}
	fmt.Println("drivers: ", drivers)
}

func (c Controller) UpdateDriver() {
	//idStr := ""
	//fmt.Println("input id: ")
	//fmt.Scan(&idStr)
	//
	//id,err := uuid.Parse(idStr)
	//if err != nil {
	//	fmt.Println("while parsing",err.Error())
	//	return
	//}
	driver := getDriverInfo()

	if err := c.Store.DriverStorage.Update(driver); err != nil {
		fmt.Println("error is while updating ", err.Error())
		return
	}
	fmt.Println("driver updated!")
}

func (c Controller) DeleteDriver() {
	idStr := ""
	fmt.Print("input id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error is while parsing", err.Error())
		return
	}
	if err = c.Store.DriverStorage.Delete(id); err != nil {
		fmt.Println("error is while deleting", err.Error())
		return
	}
	fmt.Println("driver deleted")
}

func getDriverInfo() models.Driver {
	var (
		idStr, fname, lname, phone string
		cmd                        int
	)
a:
	fmt.Print(`enter command:
				1 - create driver
				2 - update driver
`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("input id: ")
		fmt.Scan(&idStr)

		fmt.Print("input full name: ")
		fmt.Scan(&fname, &lname)

		fmt.Print("input phone: ")
		fmt.Scan(&phone)
	} else if cmd == 1 {
		fmt.Print("input full name: ")
		fmt.Scan(&fname, &lname)

		fmt.Print("input phone: ")
		fmt.Scan(&phone)
	} else {
		fmt.Println("not found")
		goto a
	}

	fullname := fname + " " + lname

	if idStr != "" {
		return models.Driver{
			ID:       uuid.MustParse(idStr),
			FullName: fullname,
			Phone:    phone,
		}
	}

	return models.Driver{
		FullName: fullname,
		Phone:    phone,
	}

}
