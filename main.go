package main

import (
	_ "vector_demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"vector_demo/internal/cmd"
)

func main() {

	cmd.Main.Run(gctx.GetInitCtx())
}
