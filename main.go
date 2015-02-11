package main

import (
	"fmt"
	"github.com/getgauge/screenshot/capture"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the filename to save the captured screenshot.")
		os.Exit(0)
	}
	err := screenshot.CaptureScreen(os.Args[1])
	if err != nil {
		panic(err)
	}
}
