package repository

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "John Doe", City: "New York", ZipCode: "10001", DateOfBirth: "2000-01-01", Status: 1},
		{CustomerID: 1002, Name: "Jane Doe", City: "New York", ZipCode: "10002", DateOfBirth: "2000-01-02", Status: 1},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, nil
}
