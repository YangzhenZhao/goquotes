package main

import (
	"fmt"
	"log"

	"github.com/YangzhenZhao/goquotes/cmd"
	"github.com/YangzhenZhao/goquotes/quotes/consts"
)

func main() {
	fmt.Println("hello")
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
	println(consts.SINA_BASE_URL)
}
