package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义结构体映射
type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Books struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Pice   float64 `db:"price"`
}

func main() {
	dsn := "root:Llh..123@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	//关闭连接
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("无法ping通过数据库: %v", err)
	}
	fmt.Println("数据连接成功！")

	//查询所有部门为 "技术部" 的员工信息
	var testEmployee []Employee
	department := "技术部"
	query1 := "select * from employees where department = ?"

	err = db.Select(&testEmployee, query1, department)
	if err != nil {
		log.Fatalf("查询数据失败！ %v", err)
	}
	//判断是否查询到数据
	if len(testEmployee) == 0 {
		log.Fatalf("未查询到条件为：%v的数据", department)
	} else {
		//遍历查询到的数据
		for _, v := range testEmployee {
			fmt.Println(v)
		}
	}

	//查询价格大于 50 元的书籍
	var books []Books
	price := 50
	query2 := "select * from books where price > ?"

	err = db.Select(&books, query2, price)
	if err != nil {
		log.Fatalf("查询数据失败！ %v", err)

	}
	//判断是否查询到数据
	if len(books) == 0 {
		log.Fatalf("未查询到查询价格大于%v的书籍", price)
	} else {
		//遍历查询到的数据
		for _, v := range books {
			fmt.Println(v)
		}
	}

}
