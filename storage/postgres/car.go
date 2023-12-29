package postgres

import (
	"carProject/models"
	"database/sql"
	"github.com/google/uuid"
)

type carRepo struct {
	DB *sql.DB
}

func NewCarRepo(db *sql.DB) carRepo {
	return carRepo{
		db,
	}
}

func (c carRepo) Insert(car models.Car) (string, error) {
	id := uuid.New()

	if _, err := c.DB.Exec(`insert into car(id,model,brand,year,driver_id) values ($1,$2,$3,$4,$5)`, id, car.Model, car.Brand, car.Year, car.DriverID); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c carRepo) GetByID(id uuid.UUID) (models.Car, error) {
	car := models.Car{}
	if err := c.DB.QueryRow(`select * from car where id = $1`, id).Scan(&car.ID, &car.Model, &car.Brand, &car.Year, &car.DriverID); err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c carRepo) GetList() ([]models.Car, error) {
	cars := []models.Car{}

	rows, err := c.DB.Query(`select * from car`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		car := models.Car{}

		if err = rows.Scan(&car.ID, &car.Model, &car.Brand, &car.Year, &car.DriverID); err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}
	return cars, nil
}

func (c carRepo) Update(car models.Car) error {
	if _, err := c.DB.Exec(`update car set model = $1, brand = $2, year = $3, driver_id = $4 where id = $5 `,
		&car.Model, &car.Brand, &car.Year, &car.DriverID, &car.ID); err != nil {
		return err
	}
	return nil
}

func (c carRepo) Delete(id uuid.UUID) error {
	if _, err := c.DB.Exec(`delete from car where id = $1`, id); err != nil {
		return err
	}

	return nil
}
