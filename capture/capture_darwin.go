package screenshot

import (
	"os/exec"
)

func CaptureScreen(filename string) error {
	cmd := exec.Command("screencapture", filename)
	err := cmd.Run()
	return err
}
