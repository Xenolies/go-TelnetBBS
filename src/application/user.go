package application

import (
	"TelnetBBS/src/inface"
	"TelnetBBS/src/utils"
	"net"
)

/*
用户模块
*/

type User struct {
	Conn    *net.TCPConn
	Level   uint   // 权限  0 admin , 1 user , 2 guest
	Addr    string // 地址
	ID      string // 系统自动生成的另外用途ID
	LoginID string // 系统自动生成的登录ID
	Name    string // 对外显示名字
	Pwd     string // 登录密码
}

func (u *User) GetID() string {
	return u.ID
}

// 已经登录注册的用户ID
func NewUser(conn *net.TCPConn) inface.IUser {
	SnowflakeID, err := utils.GenerateSnowflakeID()
	if err != nil {
		return nil
	}

	return &User{
		Conn:  conn,
		Level: 1,                          // 0 admin , 1 user , 2 guest
		Addr:  conn.RemoteAddr().String(), // 用户的IP地址和端口号
		ID:    SnowflakeID,                // 用户可以使用ID登录,
		Name:  "User_" + SnowflakeID,      // 用户的用户名
		Pwd:   "",                         // 密码
	}

}

// 未注册的游客
func NewGuestUser(conn *net.TCPConn) inface.IUser {
	SnowflakeID, err := utils.GenerateSnowflakeID()
	if err != nil {
		return nil
	}

	return &User{
		Conn:  conn,
		Level: 2,                          // 0 admin , 1 user , 2 guest
		Addr:  conn.RemoteAddr().String(), // 用户的IP地址和端口号
		ID:    "Guest_" + SnowflakeID,     // 用户可以使用ID登录,
		Name:  "Guest_" + SnowflakeID,     // 用户的用户名
		Pwd:   "",                         // 密码
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
