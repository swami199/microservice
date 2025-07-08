package domain

// createing an object for business code
type Customer struct {
	Id          string
	Name        string
	DateOfBirth string
	City        string
	Zipcode     string
	Status      int
}

//creating a port for service which is nothing but the an interface cantains function signatures
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}


