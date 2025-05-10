package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func getADBPath() string {
	return filepath.Join("adb", "adb.exe")
}

func runADBCommand(args ...string) error {
	var adbPath string = getADBPath()
	var cmd *exec.Cmd = exec.Command(adbPath, args...)
	var output []byte
	var err error

	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error: %v\noutput: %s", err, string(output))
	}
	return nil
}

func startADBServer() error {
	fmt.Println("Starting ADB server")
	return runADBCommand("start-server")
}

type adbStep struct {
	message string
	args    []string
}

func TakeScreenshot(localPath string) error {
	steps := []adbStep{
		{"Taking screenshot...", []string{"shell", "screencap", "-p", "/sdcard/screen.png"}},
		{"Pulling screenshot...", []string{"pull", "/sdcard/screen.png", localPath}},
		{"Cleaning up...", []string{"shell", "rm", "/sdcard/screen.png"}},
	}

	for _, step := range steps {
		fmt.Println(step.message)
		if err := runADBCommand(step.args...); err != nil {
			return err
		}
	}

	fmt.Printf("Screenshot saved to %s\n", localPath)
	return nil
}
