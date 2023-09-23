package main

import (
	"fmt"
	"hw1/employee1"
	"hw1/employee2"
	"hw1/employee3"
	"hw1/employee4"
	"hw1/employee5"
)

func main() {

	func_employee1 := employee1.NewEmployee1("Junior QA Engineer", 500000, "Saina 122")

	fmt.Println("Employee1 Initial Position:", func_employee1.GetPosition())
	fmt.Println("Employee1 Initial Salary:", func_employee1.GetSalary())
	fmt.Println("Employee1 Initial Address:", func_employee1.GetAddress())

	func_employee1.SetPosition("QA Automation")
	func_employee1.SetSalary(750000)
	func_employee1.SetAddress("Saina 55")

	fmt.Println("Current position:", func_employee1.GetPosition())
	fmt.Println("Current salary:", func_employee1.GetSalary())
	fmt.Println("Current address:", func_employee1.GetAddress())

	func_employee2 := employee2.NewEmployee2("Middle Software Developer", 400000, "Raimbek 144")

	fmt.Println("Employee2 Initial Position:", func_employee2.GetPosition())
	fmt.Println("Employee2 Initial Salary:", func_employee2.GetSalary())
	fmt.Println("Employee2 Initial Address:", func_employee2.GetAddress())

	func_employee2.SetPosition("Senior Software Developer")
	func_employee2.SetSalary(600000)
	func_employee2.SetAddress("Kairbekova 43")

	fmt.Println("Current position:", func_employee2.GetPosition())
	fmt.Println("Current salary:", func_employee2.GetSalary())
	fmt.Println("Current address:", func_employee2.GetAddress())

	func_employee3 := employee3.NewEmployee3("Junior Python Developer", 200000, "Auezova 40")

	fmt.Println("Employee3 Initial Position:", func_employee3.GetPosition())
	fmt.Println("Employee3 Initial Salary:", func_employee3.GetSalary())
	fmt.Println("Employee3 Initial Address:", func_employee3.GetAddress())

	func_employee3.SetPosition("Middle Python Developer")
	func_employee3.SetSalary(280000)
	func_employee3.SetAddress("Abylai khana 67")

	fmt.Println("Current position:", func_employee3.GetPosition())
	fmt.Println("Current salary:", func_employee3.GetSalary())
	fmt.Println("Current address:", func_employee3.GetAddress())

	func_employee4 := employee4.NewEmployee4("System Administrator", 450000, "Al Farabi 100")

	fmt.Println("Employee4 Initial Position:", func_employee4.GetPosition())
	fmt.Println("Employee4 Initial Salary:", func_employee4.GetSalary())
	fmt.Println("Employee4 Initial Address:", func_employee4.GetAddress())

	func_employee4.SetPosition("Team-lead")
	func_employee4.SetSalary(600000)
	func_employee4.SetAddress("Al Farabi 102")

	fmt.Println("Current position:", func_employee4.GetPosition())
	fmt.Println("Current salary:", func_employee4.GetSalary())
	fmt.Println("Current address:", func_employee4.GetAddress())

	func_employee5 := employee5.NewEmployee5("Graphic Designer", 400000, "Gogolya 17")

	fmt.Println("Employee5 Initial Position:", func_employee5.GetPosition())
	fmt.Println("Employee5 Initial Salary:", func_employee5.GetSalary())
	fmt.Println("Employee5 Initial Address:", func_employee5.GetAddress())

	func_employee5.SetPosition("Director of bootcamp")
	func_employee5.SetSalary(900000)
	func_employee5.SetAddress("Seifullina 402")

	fmt.Println("Current position:", func_employee5.GetPosition())
	fmt.Println("Current salary:", func_employee5.GetSalary())
	fmt.Println("Current address:", func_employee5.GetAddress())
}
