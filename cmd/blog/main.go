package main

import (
	"blog/internal/blog"
	"os"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := blog.NewBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
