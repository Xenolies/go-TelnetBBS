package application

import (
	"TelnetBBS/src/inface"
	"net"
)

/*
用户模块
*/

type User struct {
	Conn    *net.TCPConn
	Level   uint   // 权限  0 admin , 1 user , 2 guest
	Addr    string // 地址
	ID      string // 系统自动生成的ID
	Name    string // 对外显示名字
	LoginID string // 登录ID
	Pwd     string // 登录密码
}

func (u *User) GetID() string {
	return u.ID
}

func NewUser(conn *net.TCPConn) inface.IUser {
	return &User{
		Conn:    conn,
		Level:   1, // 0 admin , 1 user , 2 guest
		Addr:    conn.RemoteAddr().String(),
		LoginID: "",
		Name:    "",
		Pwd:     "",
	}

}

func (u *User) GetLevel() uint {
	return u.Level
}

func (u *User) SetLevel(LevenNum uint) {
	u.Level = LevenNum
}

func (u *User) GetAddr() string {
	return u.Addr
}
func (u *User) SetAddr(conn *net.TCPConn) {
	u.Conn = conn
}

func (u *User) GetLoginID() string {
	return u.LoginID
}

func (u *User) SetLoginID(loginID string) {
	u.LoginID = loginID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(Name string) {
	u.Name = Name
}

func (u *User) GetPwd() string {
	return u.Pwd
}

func (u *User) SetPwd(pwd string) {
	u.Pwd = pwd
}
