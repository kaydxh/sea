// Package application 应用层，聚合所有 Command Handler。
package application

// Application 应用层顶层结构，聚合所有 Command Handler。
type Application struct {
	Commands Commands
}

// Commands 聚合所有业务 Handler。
type Commands struct {
	SeaDateHandler SeaDateHandler
}