package main

import (
	"blog/internal/blog"
	"os"
)

func main() {
	command := blog.NewBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
