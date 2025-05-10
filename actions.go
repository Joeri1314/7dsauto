package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

func drawAction() error {
	adbPath := filepath.Join("adb", "adb.exe")
	var output []byte
	var err error

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			cmd := exec.Command(adbPath, "shell", "input", "tap", fmt.Sprintf("%d", 580), fmt.Sprintf("%d", 1640))
			output, err = cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("tap %d failed: %v\noutput: %s", i+1, err, string(output))
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func banSkip() error {
	adbPath := filepath.Join("adb", "adb.exe")
	var output []byte
	var err error

	clickPoints := [][2]int{
		{190, 1150}, {440, 1150}, {700, 1150}, {950, 1150},
		{190, 1400}, {440, 1400}, {700, 1400}, {950, 1400},
		{190, 1650}, {440, 1650}, {700, 1650}, {950, 1650},
	}

	for i, point := range clickPoints {
		cmd := exec.Command(adbPath, "shell", "input", "tap",
			fmt.Sprintf("%d", point[0]),
			fmt.Sprintf("%d", point[1]),
		)

		output, err = cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("tap %d at (%d, %d) failed: %v\noutput: %s",
				i+1, point[0], point[1], err, string(output))
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func osloSkip() error {
	adbPath := filepath.Join("adb", "adb.exe")
	var output []byte
	var err error

	startX, startY := 600, 600
	endX, endY := 600, 1300
	duration := 300 // in millis

	cmd := exec.Command(adbPath, "shell", "input", "swipe",
		fmt.Sprintf("%d", startX),
		fmt.Sprintf("%d", startY),
		fmt.Sprintf("%d", endX),
		fmt.Sprintf("%d", endY),
		fmt.Sprintf("%d", duration),
	)

	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("swipe failed: %v\noutput: %s", err, string(output))
	}

	return nil
}


func offerSkip() error {
	adbPath := filepath.Join("adb", "adb.exe")
	var output []byte
	var err error
		cmd := exec.Command(adbPath, "shell", "input", "tap", fmt.Sprintf("%d", 575), fmt.Sprintf("%d", 1830))
		output, err = cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("tap %d failed: %s", err, string(output))
		}
		return nil
}
