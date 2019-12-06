package main

import (
	"fmt"
	"os"
)

func main() {

	// ./bot archive wg c mu
	// ./bot thread c 32523534

	lenArgs := len(os.Args)
	var usedMode mode

	if lenArgs <= 1 {
		fmt.Println("Plese provide command line arguments to specify the threads to download")
		return
	} else if os.Args[1] == "archive" {
		if lenArgs > 2 {
			usedMode = Archive
		} else {
			fmt.Println("Please specify at least one board to download the archive from.")
			return
		}
	} else if os.Args[1] == "thread" {
		if lenArgs > 3 {
			usedMode = SingleThread
		} else {
			fmt.Println("Error: please specify the board name and the thread number")
			return
		}
	} else {
		fmt.Println(os.Args[1], "is not a valid parameter.")
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
