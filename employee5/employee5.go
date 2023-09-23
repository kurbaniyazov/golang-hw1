package employee5

type Employee5 struct {
	position string
	salary   float64
	address  string
}

func NewEmployee5(position string, salary float64, address string) *Employee5 {
	return &Employee5{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (e *Employee5) GetPosition() string {
	return e.position
}

func (e *Employee5) SetPosition(position string) {
	e.position = position
}

func (e *Employee5) GetSalary() float64 {
	return e.salary
}

func (e *Employee5) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Employee5) GetAddress() string {
	return e.address
}

func (e *Employee5) SetAddress(address string) {
	e.address = address
}
