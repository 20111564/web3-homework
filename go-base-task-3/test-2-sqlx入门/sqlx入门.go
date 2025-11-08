package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "gorm.io/driver/mysql"
)

var db *sqlx.DB

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 使用SQL扩展库进行查询
func test1() {

	// 查询部门为 "技术部" 的员工信息
	employees := []Employee{}
	err := db.Select(&employees, "SELECT * FROM employee WHERE department = ?", "技术部")
	if err != nil {
		fmt.Println("查询错误：", err)
	}
	fmt.Println("查询结果：", employees)

	// 查询工资最高的员工信息
	highestPaidEmployee := Employee{}
	err = db.Get(&highestPaidEmployee, "SELECT * FROM employee ORDER BY salary DESC LIMIT 1")
	if err != nil {
		fmt.Println("查询错误：", err)
	}
	fmt.Println("工资最高的员工：", highestPaidEmployee)

}

/*
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。,
编写Go代码，使用Sqlx执行一个复杂的查询，
例如查询价格大于 50 元的书籍，
并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type books struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func test2() {
	var bookList []books
	//【类型安全】意思是：校验数据类型是否匹配
	query := "SELECT * FROM books WHERE price > ?"
	price := 50.0
	err := db.Select(&bookList, query, price)
	if err != nil {
		fmt.Println("查询错误：", err)
	}
	fmt.Println("查询结果：", bookList)
}

func main() {
	fmt.Println("sqlx入门")
	//初始化数据库
	initDB()
	//题目1：使用SQL扩展库进行查询
	test1()
	//题目2：实现类型安全映射
	test2()

}

func initDB() {
	var err error
	fmt.Println("Connecting to the database...")
	dsn := "root:123456@tcp(127.0.0.1:3308)/web3-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("数据库连接成功：", db)
}
