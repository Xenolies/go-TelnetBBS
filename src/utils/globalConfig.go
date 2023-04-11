package utils

import (
	"encoding/json"
	"os"
)

/*
全局配置
*/

//type GlobalConfig struct {
//	Name             string
//	IP               string
//	Port             string
//	MaxConn          int
//	MaxPackageSize   uint32
//	WorkerPoolSize   uint32
//	MaxWorkerTaskLen uint32
//	DatabaseDriver   string
//	DatabaseName     string
//	DatabaseUserName string
//	DatabaseUserPwd  string
//}

type GlobalConfig struct {
	Name             string `json:"Name"`
	IP               string `json:"IP"`
	Port             string `json:"Port"`
	MaxConn          int    `json:"MaxConn "`
	MaxPackageSize   uint32 `json:"MaxPackageSize"`
	WorkerPoolSize   uint32 `json:"WorkerPoolSize"`
	MaxWorkerTaskLen uint32 `json:"MaxWorkerTaskLen"`
	DatabaseDriver   string `json:"DatabaseDriver"`
	DatabaseName     string `json:"DatabaseName"`
	DatabaseUserName string `json:"DatabaseUserName"`
	DatabaseUserPwd  string `json:"DatabaseUserPwd"`
	DatabasePort     int    `json:"DatabasePort"`
	DatabaseEncoding string `json:"DatabaseEncoding"`
}

var Gc *GlobalConfig // 全局配置变量 Gc

func init() {
	Gc = &GlobalConfig{
		Name:             "TelnetBBS",
		IP:               "127.0.0.1",
		Port:             "8899",
		MaxConn:          20,
		MaxPackageSize:   512,
		WorkerPoolSize:   20,
		MaxWorkerTaskLen: 20,
	}

	Gc.LoadConfig()

}

func (gc *GlobalConfig) LoadConfig() {
	data, err := os.ReadFile("conf/bbs_conf.json")
	if err != nil {
		//panic(err)
	}

	// 将 JSON 文件解析到 GlobalObj
	err = json.Unmarshal(data, &Gc)
	if err != nil {
		//panic(err)
	}
}
