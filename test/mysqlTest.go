package main

//import "database/sql"

/*
import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"
type SysUser struct {
	id     int64 `db:"id"`
	username string `db:"username"`
	password string `db:"password"`
	email string `db:"email"`
}
// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/mybasedb?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, username, password,email from sys_user where id=?"
	var u SysUser
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password,&u.email)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s password:%s email:%s\n", u.id, u.username, u.password,u.email)
}
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, username, password,email from sys_user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u SysUser
		err := rows.Scan(&u.id, &u.username, &u.password,&u.email)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s password:%s email:%s\n", u.id, u.username, u.password,u.email)
	}
}
// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, username, password,email from sys_user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u SysUser
		err := rows.Scan(&u.id, &u.username, &u.password,&u.email)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s password:%s email:%s\n", u.id, u.username, u.password,u.email)
	}
}
func main() {
	initDB()
	//queryRowDemo()
	//queryMultiRowDemo()
	prepareQueryDemo()
}
*/