package domain

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil

}
func createCustomers() CustomerRepositoryStub {
	return CustomerRepositoryStub{[]Customer{
		Customer{Id: "1001", Name: "swami", DateOfBirth: "10-03-1998", City: "Parlakhemundi", Zipcode: "761200", Status: 1},
		Customer{Id: "1002", Name: "varun", DateOfBirth: "10-03-1998", City: "Parlakhemundi", Zipcode: "761200", Status: 1},
	}}
}
