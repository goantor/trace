package trace

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var (
	node *snowflake.Node
)

func Entity() snowflake.ID {
	fmt.Println(time.Now().UnixMilli(), time.Now().UnixNano()/1e6)
	//snowflake.Epoch = time.Now().UnixMilli()
	snowflake.NodeBits = 4
	node, _ = snowflake.NewNode(1)
	return node.Generate()
}
