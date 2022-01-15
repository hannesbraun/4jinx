package main

import (
	"fmt"
	"github.com/hannesbraun/4jinx/fourchan"
	"github.com/hannesbraun/4jinx/util"
	"os"
)

func main() {

	// ./bot archive wg c mu
	// ./bot thread c 32523534

	lenArgs := len(os.Args)
	var usedMode mode

	if lenArgs <= 1 {
		fmt.Printf(util.RedColor, "Please provide command line arguments to specify the threads to download\n")
		return
	} else if os.Args[1] == "archive" {
		if lenArgs > 2 {
			usedMode = Archive
		} else {
			fmt.Printf(util.RedColor, "Please specify at least one board to download the archive from.\n")
			return
		}
	} else if os.Args[1] == "thread" {
		if lenArgs > 3 {
			usedMode = SingleThread
		} else {
			fmt.Printf(util.RedColor, "Error: please specify the board name and the thread number\n")
			return
		}
	} else {
		fmt.Println(util.RedColor + os.Args[1] + " is not a valid parameter." + util.ResetColorRaw)
		return
	}

	// Creating the root directory (if not existing yet) and chdir into it
	rootDir := "4chan_archive"
	util.Mkdir(rootDir)
	os.Chdir(rootDir)

	if usedMode == Archive {

		for boardNameIndex := 2; boardNameIndex < lenArgs; boardNameIndex++ {
			// Getting the thread numbers of the archive
			boardName := os.Args[boardNameIndex]
			threadNumbers := fourchan.GetArchiveThreadNumbers(boardName)
			util.Mkdir(boardName)

			if threadNumbers != nil {
				// Downloading every thread's images
				for threadNumber := threadNumbers.Front(); threadNumber != nil; threadNumber = threadNumber.Next() {
					fourchan.DownloadThread(boardName, threadNumber.Value.(string))
				}
			}
		}

	} else if usedMode == SingleThread {

		// Getting the boardName and the threadNumber and creating its directory
		boardName := os.Args[2]
		util.Mkdir(boardName)
		threadNumber := os.Args[3]

		// Downloading the thread
		fourchan.DownloadThread(boardName, threadNumber)

	}

}

type mode int

const (
	// SingleThread mode: downloading a single thread
	SingleThread = iota
	// Archive mode: downloading the complete archive
	Archive = iota
)
