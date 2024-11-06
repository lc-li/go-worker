package main

import (
	"fmt"
	"log"

	"go-worker/i18n"
)

func main() {
	// 创建翻译器实例，设置默认语言为英语
	translator := i18n.New("en")

	// 加载翻译文件
	err := translator.LoadTranslationFile("en", "locales/en.json")
	if err != nil {
		log.Fatal(err)
	}

	err = translator.LoadTranslationFile("zh", "locales/zh.json")
	if err != nil {
		log.Fatal(err)
	}

	// 测试英文翻译
	fmt.Println(translator.T("en", "welcome", "John"))
	fmt.Println(translator.T("en", "goodbye"))
	fmt.Println(translator.T("en", "hello_world"))

	// 测试中文翻译
	fmt.Println(translator.T("zh", "welcome", "张三"))
	fmt.Println(translator.T("zh", "goodbye"))
	fmt.Println(translator.T("zh", "hello_world"))
}
