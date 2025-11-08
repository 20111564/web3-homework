package main

import "fmt"

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func shape() {
	rectangle := &Rectangle{Width: 5, Height: 10}
	circle := &Circle{Radius: 7}

	fmt.Println("矩形面积 =", rectangle.Area())
	fmt.Println("矩形周长 =", rectangle.Perimeter())
	fmt.Println("圆面积 =", circle.Area())
	fmt.Println("圆周长 =", circle.Perimeter())
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，
包含 Name 和 Age 字段，
再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	println("Name:", e.Name)
	println("Age:", e.Age)
	println("EmployeeID:", e.EmployeeID)
}
func employeeInfo() {
	employee := &Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
	}
	employee.PrintInfo()
}
func main() {
	//题目一，形状
	shape()
	//题目二，员工信息
	employeeInfo()
}
