package employee4

type Employee4 struct {
	position string
	salary   float64
	address  string
}

func NewEmployee4(position string, salary float64, address string) *Employee4 {
	return &Employee4{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (e *Employee4) GetPosition() string {
	return e.position
}

func (e *Employee4) SetPosition(position string) {
	e.position = position
}

func (e *Employee4) GetSalary() float64 {
	return e.salary
}

func (e *Employee4) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Employee4) GetAddress() string {
	return e.address
}

func (e *Employee4) SetAddress(address string) {
	e.address = address
}
