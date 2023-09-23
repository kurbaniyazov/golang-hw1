package employee2

type Employee2 struct {
	position string
	salary   float64
	address  string
}

func NewEmployee2(position string, salary float64, address string) *Employee2 {
	return &Employee2{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (e *Employee2) GetPosition() string {
	return e.position
}

func (e *Employee2) SetPosition(position string) {
	e.position = position
}

func (e *Employee2) GetSalary() float64 {
	return e.salary
}

func (e *Employee2) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Employee2) GetAddress() string {
	return e.address
}

func (e *Employee2) SetAddress(address string) {
	e.address = address
}
