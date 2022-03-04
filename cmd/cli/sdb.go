package main

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/yemingfeng/sdb/internal/cli"
)

func main() {
	shell := ishell.New()
	shell.Println("sdb cli")
	cli.RegisterBitsetCmd(shell)
	shell.Run()
}
