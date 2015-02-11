package main

import (
	"fmt"
	"github.com/getgauge/screenshot/capture"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <target filepath>\n", os.Args[0])
		os.Exit(0)
	}
	err := capture.CaptureScreen(os.Args[1])
	if err != nil {
		panic(err)
	}
}
