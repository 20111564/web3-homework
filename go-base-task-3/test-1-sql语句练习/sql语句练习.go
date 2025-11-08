package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Student struct {
	ID    int32
	Name  string
	Age   int32
	Grade string
}

// 基本CRUD操作
func curd() {
	var db = initDB()
	//插入数据
	insertStudents := []Student{
		{Name: "张三", Age: 20, Grade: "三年级"},
		{Name: "李四", Age: 18, Grade: "二年级"},
	}
	result := db.Create(&insertStudents)
	fmt.Println("影响行数：", result.RowsAffected)
	fmt.Println("插入信息：", insertStudents)
	//查询数据
	var searchStudents []Student
	db.Debug().Where("age > ?", 18).Find(&searchStudents)
	fmt.Println("查询结果：", searchStudents)
	//跟新数据
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	//删除数据
	db.Where("age < ?", 15).Delete(&Student{})
}

/*
假设有两个表：
accounts 表（包含字段 id 主键， balance 账户余额）
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务。
*/
var mu = sync.Mutex{}

type Account struct {
	ID      uint32
	Balance float32
}
type Transaction struct {
	ID            uint32
	FromAccountID uint32
	ToAccountID   uint32
	Amount        float32
}

func transaction(_fromAccountID, _toAccountID uint32, _amount float32) {
	defer mu.Unlock()
	db := initDB()
	mu.Lock()
	err := db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		resultFromAccount := tx.First(&fromAccount, _fromAccountID)
		resultToAccount := tx.First(&toAccount, _toAccountID)
		if (resultFromAccount.Error != nil) || (resultToAccount.Error != nil) {
			return fmt.Errorf("account not found")
		}
		if fromAccount.Balance < _amount {
			return fmt.Errorf("fromAccount.Balance < _amount")
		}
		fromAccount.Balance -= _amount
		toAccount.Balance += _amount

		newTransaction := Transaction{}
		newTransaction.FromAccountID = _fromAccountID
		newTransaction.ToAccountID = _toAccountID
		newTransaction.Amount = _amount

		//保存或者跟新数据
		tx.Save(&fromAccount)
		tx.Save(&toAccount)
		tx.Create(&newTransaction)
		return nil
	})
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction succeeded")
	}
}

func transactionTest() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			transaction(1, 2, 100)
		}()
	}
	wg.Wait()
}
func main() {
	// 题目一：基本crud操作
	//curd()
	// 题目二：事务语句
	transactionTest()
}

func initDB() *gorm.DB {
	fmt.Println("Connecting to the database...")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3308)/web3-test?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
