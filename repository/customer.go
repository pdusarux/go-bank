package repository

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)    // select *
	GetById(int) (*Customer, error) // where id limit 1
}
