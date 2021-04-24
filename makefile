SOURCECODE = main.go downloadThread.go getArchiveThreadNumbers.go getImageURLs.go tools.go
MODULE = go.mod go.sum

all: $(SOURCECODE) $(MODULE)
	mkdir -p bin
	go build -o bin/4jinx .

clean:
	rm -rf bin/
	rm -rf dist/

dist: $(SOURCECODE) $(MODULE)
	mkdir -p dist
	env GOOS=darwin GOARCH=amd64 go build -o dist/4jinx-darwin-amd64
	env GOOS=freebsd GOARCH=amd64 go build -o dist/4jinx-freebsd-amd64
	env GOOS=freebsd GOARCH=386 go build -o dist/4jinx-freebsd-386
	env GOOS=freebsd GOARCH=arm go build -o dist/4jinx-freebsd-arm
	env GOOS=linux GOARCH=386 go build -o dist/4jinx-linux-386
	env GOOS=linux GOARCH=amd64 go build -o dist/4jinx-linux-amd64
	env GOOS=linux GOARCH=arm go build -o dist/4jinx-linux-arm
	env GOOS=linux GOARCH=arm64 go build -o dist/4jinx-linux-arm64
	env GOOS=linux GOARCH=ppc64 go build -o dist/4jinx-linux-ppc64
	env GOOS=linux GOARCH=ppc64le go build -o dist/4jinx-linux-ppc64le
	env GOOS=linux GOARCH=mips go build -o dist/4jinx-linux-mips
	env GOOS=linux GOARCH=mipsle go build -o dist/4jinx-linux-mipsle
	env GOOS=linux GOARCH=mips64 go build -o dist/4jinx-linux-mips64
	env GOOS=linux GOARCH=mips64le go build -o dist/4jinx-linux-mips64le
	env GOOS=netbsd GOARCH=amd64 go build -o dist/4jinx-netbsd-amd64
	env GOOS=netbsd GOARCH=386 go build -o dist/4jinx-netbsd-386
	env GOOS=netbsd GOARCH=arm go build -o dist/4jinx-netbsd-arm
	env GOOS=openbsd GOARCH=amd64 go build -o dist/4jinx-openbsd-amd64
	env GOOS=openbsd GOARCH=386 go build -o dist/4jinx-openbsd-386
	env GOOS=openbsd GOARCH=arm go build -o dist/4jinx-openbsd-arm
	env GOOS=solaris GOARCH=amd64 go build -o dist/4jinx-solaris-amd64
	env GOOS=windows GOARCH=amd64 go build -o dist/4jinx-windows-amd64.exe
	env GOOS=windows GOARCH=386 go build -o dist/4jinx-windows-386.exe

.PHONY: all clean dist
