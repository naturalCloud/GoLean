package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/88250/lute"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// markdown 格式化引擎
func main() {

	filePathStr := flag.String("p", "", "文件路径")

	flag.Parse()

	filePath := *filePathStr

	_, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		log.Fatalf("%s 文件不存在！", filePath)
	}

	engine := lute.New()

	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`^---([\d\D]*)---`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	file, _ := os.Open(filePath)

	defer file.Close()
	all, _ := io.ReadAll(file)
	originStr := string(all)

	result1 := reg1.ReplaceAllStringFunc(originStr, func(s string) string {
		return ""
	})

	builder := strings.Builder{}
	builder.WriteString(reg1.FindString(originStr))

	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	formatStr := engine.FormatStr("demo", result1)

	builder.WriteString(formatStr)

	ioutil.WriteFile(filePath, []byte(builder.String()), 0666)

}
