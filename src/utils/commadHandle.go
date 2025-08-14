package utils

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
	s = strings.TrimSpace(s) // 先去掉前后空格
	if len(s) == 0 {
		return "", "", false
	}

	if s[0] != '@' {
		return "", s, false
	}

	idx := strings.IndexByte(s, cut)
	if idx == -1 { // 没有空格，只有命令
		return s, "", true
	}

	cmd := s[:idx]
	args := strings.TrimSpace(s[idx+1:])
	return cmd, args, true
}
