package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (r CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return r.customers, nil
}

func NewCostumerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "666", Name: "John", City: "RJ", Zipcode: "232323", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "999", Name: "Joane", City: "SP", Zipcode: "232323", DateofBirth: "2000-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
