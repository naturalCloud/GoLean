package main

import (
	"fmt"
	"github.com/88250/lute"
)

// markdown 格式化引擎
func main() {
	engine := lute.New()
	originStr := `
|col1|col2  |       col3   |
---           |---------------|--
col1 without left pipe      |   this is col2   | col3 without right pipe
                                 ||need align cell|
`
	formatStr := engine.FormatStr("demo", originStr)
	fmt.Println(formatStr)
}
