package main

import (
	_ "gfDeomo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gfDeomo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
