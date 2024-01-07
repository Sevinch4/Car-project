package postgres

import (
	"carProject/models"
	"database/sql"
	"github.com/google/uuid"
)

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) driverRepo {
	return driverRepo{
		DB: db,
	}
}

func (d driverRepo) Insert(driver models.Driver) (string, error) {
	id := uuid.New()
	if _, err := d.DB.Exec(`insert into driver (id,full_name,phone) values($1, $2, $3)`,
		id, driver.FullName, driver.Phone); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (d driverRepo) GetByID(id string) (models.Driver, error) {
	driver := models.Driver{}
	if err := d.DB.QueryRow(`select * from driver where id = $1`, id).Scan(&driver.ID,
		&driver.FullName, &driver.Phone); err != nil {
		return models.Driver{}, err
	}
	return driver, nil
}

func (d driverRepo) GetList() ([]models.Driver, error) {
	drivers := []models.Driver{}

	rows, err := d.DB.Query(`select * from driver`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		driver := models.Driver{}
		if err = rows.Scan(&driver.ID, &driver.FullName, &driver.Phone); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (d driverRepo) Update(driver models.Driver) error {
	if _, err := d.DB.Exec(`update driver set full_name = $1,phone = $2 where id = $3`,
		&driver.FullName, &driver.Phone, &driver.ID); err != nil {
		return err
	}
	return nil
}

func (d driverRepo) Delete(id string) error {
	if _, err := d.DB.Exec(`delete from driver where id = $1`, id); err != nil {
		return err
	}
	return nil
}
