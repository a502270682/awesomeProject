package strategy

import "fmt"

/*
策略模式
属于：行为型模式
优点
提供了对 "开闭原则"的完美支持。
避免使用多重条件转移语句。
提供了管理相关的算法族的办法。
缺点：
客户端必须知道所有的策略类。
策略模式将造成产生很多策略类。
适合场景
需要动态地在几种算法中选择一种。
多个类区别仅在于它们的行为或算法不同的场景。
*/

// 实现一个日志记录器，相当于策略的上下文
type LoggerManager struct {
	Logger
}

// 抽象的日志
type Logger interface {
	Info(msg string)
	Error(msg string)
}

func NewLoggerManager(logger Logger) *LoggerManager {
	return &LoggerManager{
		logger,
	}
}

// 实现具体的日志： 文件方式
type FileLogger struct {
}

func (f *FileLogger) Info(msg string) {
	fmt.Println("文件info输出", msg)
}

func (f *FileLogger) Error(msg string) {
	fmt.Println("文件error输出", msg)
}

// 实现具体的日志： 数据库方式
type DbLogger struct {
}

func (f *DbLogger) Info(msg string) {
	fmt.Println("数据库info输出", msg)
}

func (f *DbLogger) Error(msg string) {
	fmt.Println("数据库error输出", msg)
}
