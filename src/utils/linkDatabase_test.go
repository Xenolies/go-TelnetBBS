package utils

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestNewlineDatabase(t *testing.T) {

	DSN := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=%s", "root", "123456", 3306, "go_project", "utf8")
	db, err := sql.Open("mysql", DSN)
	if err != nil { //不会校验用户名密码是否正确，只校验数据源格式
		fmt.Printf("sql.Open err: %v", err) //dsn格式不正确时报错
	}
	db.Close()
}
