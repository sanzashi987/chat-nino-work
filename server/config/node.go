package config

import "github.com/bwmarrin/snowflake"

var SnowflakeNode *snowflake.Node

func CreateSnowflakeNode() {
	var err error
	SnowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		panic("Error generate snowflake node")
	}
}
