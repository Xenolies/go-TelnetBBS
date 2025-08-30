package main

import (
	"TelnetBBS/routers"
	"TelnetBBS/src/application"
	"TelnetBBS/src/inface"
)

func main() {
	s := application.NewServer()

	s.AddRouter("@login", &routers.LoginRouter{})

	s.Serve()

}

type PageIndex struct {
	inface.IPage
}
