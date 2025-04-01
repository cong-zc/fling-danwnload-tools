package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

func U_抓取网页数据(搜索关键词 string) map[string]string {

	// 创建上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置超时
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 目标网页
	baseUrl := "https://flingtrainer.com/"
	params := url.Values{}
	params.Set("s", 搜索关键词) // 设置查询参数
	url := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	fmt.Println("请求的URL:", url)

	// 等待页面中某个元素出现（例如选择器为 ".content" 的元素）
	var href string
	var text string // 定义存储结构
	// var attributes []map[string]string // 定义存储结构

	// 执行任务列表
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("article a", chromedp.ByQuery), // 根据实际页面调整选择器
		// 页面加载完成后，提取"article a"元素的 href 属性和文本内容
		chromedp.AttributeValue("article a", "href", &href, nil, chromedp.ByQuery),
		chromedp.Text("article a", &text, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("href:", href)
	返回值 := 获取详情页(href) // 调用获取详情页函数
	return 返回值         // 返回结果

}

func 获取详情页(href string) map[string]string {
	// 创建上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置超时
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var text string        // 定义存储结构
	var downloadUrl string // 定义存储结构
	// 执行任务列表
	err := chromedp.Run(ctx,
		chromedp.Navigate(href),
		chromedp.WaitVisible("div.content", chromedp.ByQuery), // 根据实际页面调整选择器
		chromedp.Text("h1.post-title", &text, chromedp.NodeVisible, chromedp.ByQuery),
		// 提取下载链接
		chromedp.AttributeValue(".download-attachments a", "href", &downloadUrl, nil, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("详情页标题:", text)
	fmt.Println("下载链接:", downloadUrl)
	return map[string]string{
		"标题":   text,
		"下载链接": downloadUrl,
	}
}
