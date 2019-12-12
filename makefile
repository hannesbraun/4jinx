all: main.go downloadThread.go getArchiveThreadNumbers.go getImageURLs.go tools.go
	mkdir -p bin
	go build -o bin/4jinx main.go downloadThread.go getArchiveThreadNumbers.go getImageURLs.go tools.go

clean:
	rm -r bin/
