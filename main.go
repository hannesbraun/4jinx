package main

import (
	"fmt"
	"os"
)

const (
	BlueColor = "\033[1;34m%s\033[0;0m"
	BlueColorRaw = "\033[1;34m"
	RedColor = "\033[1;31m%s\033[0;0m"
	RedColorRaw = "\033[1;31m"
	ResetColorRaw = "\033[0;0m"
	YellowColor = "\033[0;33m%s\033[0;0m"
)

func main() {

	// ./bot archive wg c mu
	// ./bot thread c 32523534

	lenArgs := len(os.Args)
	var usedMode mode

	if lenArgs <= 1 {
		fmt.Printf(RedColor, "Please provide command line arguments to specify the threads to download\n")
		return
	} else if os.Args[1] == "archive" {
		if lenArgs > 2 {
			usedMode = Archive
		} else {
			fmt.Printf(RedColor, "Please specify at least one board to download the archive from.\n")
			return
		}
	} else if os.Args[1] == "thread" {
		if lenArgs > 3 {
			usedMode = SingleThread
		} else {
			fmt.Printf(RedColor, "Error: please specify the board name and the thread number\n")
			return
		}
	} else {
		fmt.Println(RedColor + os.Args[1] + " is not a valid parameter." + ResetColorRaw)
		return
	}

	// Creating the root direcotry (if not existing yet) and chdir into it
	rootDir := "4chan_archive"
	mkdir(rootDir)
	os.Chdir(rootDir)

	if usedMode == Archive {

		for boardNameIndex := 2; boardNameIndex < lenArgs; boardNameIndex++ {
			// Geting the thread numbers of the archive
			boardName := os.Args[boardNameIndex]
			threadNumbers := getArchiveThreadNumbers(boardName)
			mkdir(boardName)

			if threadNumbers != nil {
				// Downloading every thread's images
				for threadNumber := threadNumbers.Front(); threadNumber != nil; threadNumber = threadNumber.Next() {
					downloadThread(boardName, threadNumber.Value.(string))
				}
			}
		}

	} else if usedMode == SingleThread {

		// Geting the boardName and the trheadNumber and creating its directory
		boardName := os.Args[2]
		mkdir(boardName)
		threadNumber := os.Args[3]

		// Downloading the thread
		downloadThread(boardName, threadNumber)

	}

}

type mode int

const (
	// SingleThread mode: downloading a single thread
	SingleThread = iota
	// Archive mode: downloading the complete archive
	Archive = iota
)
