package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.topgoer.cn/" // 替换为你要爬取的网页URL

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求网页时发生错误:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网页内容时发生错误:", err)
		return
	}

	fmt.Println(string(body))
}
