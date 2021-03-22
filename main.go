package main

import (
	"log"

	"github.com/YangzhenZhao/goquotes/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
