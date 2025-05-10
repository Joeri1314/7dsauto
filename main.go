package main

import (
	"fmt"
	"sync"
)

func main() {

	err := startADBServer()
	if err != nil {
		fmt.Println("Failed to start ADB:", err)
		return
	}

	loop()
}


func loop() {
	var err error
	var rolledAmount int = 0

	fmt.Print("Enter card amount: ")
	var energyAmount int
	_, err = fmt.Scan(&energyAmount)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	maxRolls := (energyAmount - (energyAmount %3)) / 3

	for {
		err = TakeScreenshot("screen.png")
		if err != nil {
			fmt.Println("Failed to take screenshot:", err)
			return
		}

		var wg sync.WaitGroup
		matchResults := make(chan string, 5) // buffered so goroutines won't block

		images := []string{
			"energyRollable",
			"banSkip",
			"osloSkip",
			"offerSkip",
		}

		for _, name := range images {
			wg.Add(1)
			go func(match string) {
				defer wg.Done()
				if imageRec(match) {
					matchResults <- match
				}
			}(name)
		}

		// Wait for all routines then close channel
		go func() {
			wg.Wait()
			close(matchResults)
		}()

		// Handle results
		rolled := false
		for match := range matchResults {
			if match == "energyRollable" && !rolled {
				drawAction()
				rolledAmount++
				fmt.Printf("%d / %d", rolledAmount, maxRolls)
				println()
				rolled = true
			}
			if match == "banSkip" && !rolled {
				banSkip()
				rolled = true
			}
			if match == "osloSkip" && !rolled {
				osloSkip()
				rolled = true
			}
			if match == "offerSkip" && !rolled {
				offerSkip()
				rolled = true
			}
		}
		if (rolledAmount >= maxRolls) {
			return;
		}
	}
}