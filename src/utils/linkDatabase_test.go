package utils

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestNewlineDatabase(t *testing.T) {

	DSN := fmt.Sprintf("%s:%s@tcp(192.168.50.81/:%d)/%s?charset=%s", "root", "20021121", 3306, "go_telnetBBS", "utf8")
	db, err := sql.Open("mysql", DSN)
	fmt.Println(DSN)
	fmt.Println(err)
	fmt.Println(db.Query("user"))
	if err != nil { //不会校验用户名密码是否正确，只校验数据源格式
		fmt.Printf("sql.Open err: %v", err) //dsn格式不正确时报错
	}
	db.Close()
}
