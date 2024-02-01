package main

import (
	"fmt"
	"interpreterInGo/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	// 打印欢迎信息
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
