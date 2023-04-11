package util

import (
	"strings"
)

/*
命令处理 ,分离命令和输入的数据
*/

// GetCommand 获取输入Command
func GetCommand(input string) (string, string, bool) {
	return SubStrRange(input, ' ')
}

func SubStrRange(s string, cut byte) (string, string, bool) {
	var n int
	if s[0] == '@' {
		for i := range s {
			if s[i] == cut {
				break
			}
			n++
		}
		return s[:n], strings.TrimSpace(s[n:]), true
	} else {
		return "", s, false
	}

}
