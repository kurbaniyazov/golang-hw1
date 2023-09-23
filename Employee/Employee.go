package company

type Employee interface {
  get_position() string
  set_position(string)
  get_salary() float64
  set_salary(float64)
  get_address() string
  set_address(string)
}