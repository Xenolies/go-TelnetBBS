package application

import (
	"TelnetBBS/src/inface"
)

/*
消息实现层
*/

type Message struct {
	Data       string
	Command    string
	DataLen    uint32
	MaxDataLen uint32
}

// NewMessage 创建新的消息
func NewMessage(command string, data string) inface.IMessage {
	return &Message{
		Data:    data,
		Command: command,
		DataLen: uint32(len(data)),
	}
}

func (m *Message) GetData() string {
	return m.Data
}

func (m *Message) GetCommand() string {
	return m.Command
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}
