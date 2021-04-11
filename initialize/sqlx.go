package initialize

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func Sqlx() *sqlx.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return sqlxMysql()
	case "sqlserver":
		return sqlxSqlServer()
	default:
		return sqlxMysql()
	}
}


func sqlxMysql() *sqlx.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path +")/" + m.Dbname + "?" +m.Config
	db,err :=sqlx.Connect("mysql",dsn)
	if err != nil {
		fmt.Println("connect DB failed, err:%v\n", err)
		log.Fatalln(err)
		return nil
	}

	fmt.Println("数据库链接成功"+dsn)
	db.SetMaxOpenConns(200)
	db.SetConnMaxIdleTime(10)
	return db
}

func sqlxSqlServer() *sqlx.DB {
	return nil
}
