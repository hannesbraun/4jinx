package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"sync"
)

var activeTasks sync.WaitGroup

func downloadThread(boardName string, threadNumber string) {
	if !isExisting(boardName, threadNumber) {
		// Getting all image urls
		imageURLs := getImageURLs(boardName, threadNumber)

		if imageURLs != nil {
			var activeTasksCounter = 0

			// Download every image of the thread
			for imageURL := imageURLs.Front(); imageURL != nil; imageURL = imageURL.Next() {
				if activeTasksCounter > 64 {
					fmt.Println("Cooldown: 64 files are opened")
					activeTasks.Wait()
					activeTasksCounter = 0
				}
				activeTasks.Add(1)
				activeTasksCounter = activeTasksCounter + 1
				go downloadImage(&DownloadTask{boardName, threadNumber, "http:" + imageURL.Value.(string)})
			}
			activeTasks.Wait()
		}
	}
}

func isExisting(boardName string, threadNumber string) bool {
	threadDir := boardName + "/" + threadNumber
	_, err := os.Stat(threadDir)

	if os.IsNotExist(err) {
		mkdirCommand := exec.Command("mkdir", threadDir)
		err = mkdirCommand.Run()
		if err != nil {
			fmt.Println("An error occured while creating the thread directory.")
		}
		return false
	} else if os.IsExist(err) {
		// Directory already existing, leaving
		return true
	} else {
		// Unknown Error
		fmt.Println("An error occured while creating the thread directory.")
		return true
	}
}

func downloadImage(imageTask *DownloadTask) {
	defer activeTasks.Done()

	fmt.Println("Downloading image " + imageTask.imageURL)

	request, _ := http.NewRequest("GET", imageTask.imageURL, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		filename := path.Base(request.URL.Path)
		file, err := os.Create(imageTask.boardName + "/" + imageTask.threadNumber + "/" + filename)
		if err != nil {
			log.Fatal(err)
		} else {
			defer file.Close()

			_, err := io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}

		}

	}

}

// DownloadTask is a type including the necessary information to download an image
type DownloadTask struct {
	// The short name of the 4chan board (e.g. "wg")
	boardName string

	// The number of the thread as a string
	threadNumber string

	// The actual URL
	imageURL string
}
