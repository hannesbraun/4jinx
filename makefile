SOURCECODE = main.go downloadThread.go getArchiveThreadNumbers.go getImageURLs.go tools.go

all: $(SOURCECODE)
	mkdir -p bin
	go build -o bin/4jinx $(SOURCECODE)

clean:
	rm -rf bin/
	rm -rf dist/

dist: $(SOURCECODE)
	mkdir -p dist
	env GOOS=darwin GOARCH=amd64 go build -o dist/4jinx-darwin-amd64 $(SOURCECODE)
	env GOOS=freebsd GOARCH=amd64 go build -o dist/4jinx-freebsd-amd64 $(SOURCECODE)
	env GOOS=freebsd GOARCH=386 go build -o dist/4jinx-freebsd-386 $(SOURCECODE)
	env GOOS=freebsd GOARCH=arm go build -o dist/4jinx-freebsd-arm $(SOURCECODE)
	env GOOS=linux GOARCH=386 go build -o dist/4jinx-linux-386 $(SOURCECODE)
	env GOOS=linux GOARCH=amd64 go build -o dist/4jinx-linux-amd64 $(SOURCECODE)
	env GOOS=linux GOARCH=arm go build -o dist/4jinx-linux-arm $(SOURCECODE)
	env GOOS=linux GOARCH=arm64 go build -o dist/4jinx-linux-arm64 $(SOURCECODE)
	env GOOS=linux GOARCH=ppc64 go build -o dist/4jinx-linux-ppc64 $(SOURCECODE)
	env GOOS=linux GOARCH=ppc64le go build -o dist/4jinx-linux-ppc64le $(SOURCECODE)
	env GOOS=linux GOARCH=mips go build -o dist/4jinx-linux-mips $(SOURCECODE)
	env GOOS=linux GOARCH=mipsle go build -o dist/4jinx-linux-mipsle $(SOURCECODE)
	env GOOS=linux GOARCH=mips64 go build -o dist/4jinx-linux-mips64 $(SOURCECODE)
	env GOOS=linux GOARCH=mips64le go build -o dist/4jinx-linux-mips64le $(SOURCECODE)
	env GOOS=netbsd GOARCH=amd64 go build -o dist/4jinx-netbsd-amd64 $(SOURCECODE)
	env GOOS=netbsd GOARCH=386 go build -o dist/4jinx-netbsd-386 $(SOURCECODE)
	env GOOS=netbsd GOARCH=arm go build -o dist/4jinx-netbsd-arm $(SOURCECODE)
	env GOOS=openbsd GOARCH=amd64 go build -o dist/4jinx-openbsd-amd64 $(SOURCECODE)
	env GOOS=openbsd GOARCH=386 go build -o dist/4jinx-openbsd-386 $(SOURCECODE)
	env GOOS=openbsd GOARCH=arm go build -o dist/4jinx-openbsd-arm $(SOURCECODE)
	env GOOS=solaris GOARCH=amd64 go build -o dist/4jinx-solaris-amd64 $(SOURCECODE)
	env GOOS=windows GOARCH=amd64 go build -o dist/4jinx-windows-amd64.exe $(SOURCECODE)
	env GOOS=windows GOARCH=386 go build -o dist/4jinx-windows-386.exe $(SOURCECODE)

.PHONY: all clean dist
