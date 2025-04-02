package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// 一个简单供前端调用的函数
func (a *App) G_获取游戏名称(搜索关键词 string) map[string]string {

	fmt.Println("游戏名称:", 搜索关键词)
	返回值 := U_抓取网页数据(搜索关键词)
	return 返回值
}
