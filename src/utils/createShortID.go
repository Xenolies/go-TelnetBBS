package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflakeID() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	id := node.Generate()
	return id.String(), nil
}
