package main

import (
	"fmt"
	"os"

	"github.com/aloysZY/MyOperatorProjects/toolsets/cmd/app"
)

func main() {
	if err := app.RootCmd.Execute(); err != nil {
		// 如果解析 flagset 出错，将 panic 并将 error 信息输出到 sugaredLogger
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
