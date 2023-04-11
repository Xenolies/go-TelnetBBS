package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	Driver string
	DSN    string
}

// NewDB 获取数据库DB对象
func NewDB() *sql.DB {
	DSN := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=%s", Gc.DatabaseUserName, Gc.DatabaseUserPwd, Gc.DatabasePort, Gc.DatabaseName, Gc.DatabaseEncoding)

	db, err := sql.Open(Gc.DatabaseDriver, DSN)
	if err != nil { //不会校验用户名密码是否正确，只校验数据源格式
		fmt.Printf("sql.Open err: %v", err) //dsn格式不正确时报错
	}
	return db
}
func (db *DataBase) getDatabaseObject() {

}

//driver := "mysql"
////root为数据库用户名，后面为密码，tcp代表tcp协议，test处填写自己的数据库名称
//DSN := "root:123123@tcp(127.0.0.1:3306)/godatabase?charset=utf8"
//
//db, err := sql.Open(drever, DSN)
//if err != nil { //不会校验用户名密码是否正确，只校验数据源格式
//fmt.Printf("sql.Open err: %v", err) //dsn格式不正确时报错
//return
//}
