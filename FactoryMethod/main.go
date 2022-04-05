// Builder 设计模式·创建型模式·工厂方法
package main

import (
	"flag"
	"fmt"
	"time"
)

// Logger 日志接口定义
type Logger interface {
	Write(str string) // 写入日志
}

// LogFile 文件日志
type LogFile struct{}

func (o *LogFile) Write(str string) { fmt.Println("写入到文件：", str) }

// LogNet 网络日志
type LogNet struct{}

func (o *LogNet) Write(str string) { fmt.Println("通过网络写日志：", str) }

// LogNull 丢弃日志
type LogNull struct{}

func (o *LogNull) Write(str string) { fmt.Println("什么都不做") }

// 工厂
func makeLogger(where string) Logger {
	switch where {
	case "net":
		return &LogNet{}
	case "file":
		return &LogFile{}
	default:
		return &LogNull{}
	}
}

var logWhere = flag.String("w", "", "[file|net]日志存储到哪里？")

// go run .
// go run . -w file
// go run . -w net
func main() {
	flag.Parse()
	logger := makeLogger(*logWhere)
	logger.Write(time.Now().Format("2006-01-02 15:04:05"))
}
