SOURCECODE = cmd/*.go fourchan/*.go util/*.go
MODULE = go.mod go.sum

all: bin/4jinx

bin/4jinx: $(SOURCECODE) $(MODULE)
	mkdir -p bin
	go build -o bin/4jinx cmd/4jinx.go

clean:
	rm -rf bin/
	rm -rf dist/

dist: $(SOURCECODE) $(MODULE)
	mkdir -p dist
	env GOOS=darwin GOARCH=amd64 go build -o dist/4jinx-darwin-amd64 cmd/4jinx.go
	env GOOS=freebsd GOARCH=amd64 go build -o dist/4jinx-freebsd-amd64 cmd/4jinx.go
	env GOOS=freebsd GOARCH=386 go build -o dist/4jinx-freebsd-386 cmd/4jinx.go
	env GOOS=freebsd GOARCH=arm go build -o dist/4jinx-freebsd-arm cmd/4jinx.go
	env GOOS=linux GOARCH=386 go build -o dist/4jinx-linux-386 cmd/4jinx.go
	env GOOS=linux GOARCH=amd64 go build -o dist/4jinx-linux-amd64 cmd/4jinx.go
	env GOOS=linux GOARCH=arm go build -o dist/4jinx-linux-arm cmd/4jinx.go
	env GOOS=linux GOARCH=arm64 go build -o dist/4jinx-linux-arm64 cmd/4jinx.go
	env GOOS=linux GOARCH=ppc64 go build -o dist/4jinx-linux-ppc64 cmd/4jinx.go
	env GOOS=linux GOARCH=ppc64le go build -o dist/4jinx-linux-ppc64le cmd/4jinx.go
	env GOOS=linux GOARCH=mips go build -o dist/4jinx-linux-mips cmd/4jinx.go
	env GOOS=linux GOARCH=mipsle go build -o dist/4jinx-linux-mipsle cmd/4jinx.go
	env GOOS=linux GOARCH=mips64 go build -o dist/4jinx-linux-mips64 cmd/4jinx.go
	env GOOS=linux GOARCH=mips64le go build -o dist/4jinx-linux-mips64le cmd/4jinx.go
	env GOOS=netbsd GOARCH=amd64 go build -o dist/4jinx-netbsd-amd64 cmd/4jinx.go
	env GOOS=netbsd GOARCH=386 go build -o dist/4jinx-netbsd-386 cmd/4jinx.go
	env GOOS=netbsd GOARCH=arm go build -o dist/4jinx-netbsd-arm cmd/4jinx.go
	env GOOS=openbsd GOARCH=amd64 go build -o dist/4jinx-openbsd-amd64 cmd/4jinx.go
	env GOOS=openbsd GOARCH=386 go build -o dist/4jinx-openbsd-386 cmd/4jinx.go
	env GOOS=openbsd GOARCH=arm go build -o dist/4jinx-openbsd-arm cmd/4jinx.go
	env GOOS=solaris GOARCH=amd64 go build -o dist/4jinx-solaris-amd64 cmd/4jinx.go
	env GOOS=windows GOARCH=amd64 go build -o dist/4jinx-windows-amd64.exe cmd/4jinx.go
	env GOOS=windows GOARCH=386 go build -o dist/4jinx-windows-386.exe cmd/4jinx.go

.PHONY: all clean dist
