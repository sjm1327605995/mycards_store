package snow

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

}

// GenID 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}
