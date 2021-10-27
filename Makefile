PROJECT=revproxy

.PHONY: build

build:
	# for the current OS only
	go build
build:
	# for windows-amd64 only
	env GOOS=windows GOARCH=amd64 go build -o bin/$(PROJECT)-linux-amd64
buildwin:
	# for windows-amd64 only
	env GOOS=windows GOARCH=amd64 go build -o bin/$(PROJECT)-linux-amd64
buildlinux: 
	# for linux-amd64 only
	env GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT)-linux-amd64
buildall:
	# for all platforms
	env GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT)-linux-amd64
	env GOOS=darwin GOARCH=amd64 go build -o bin/$(PROJECT)-darwin-amd64
	env GOOS=darwin GOARCH=arm64 go build -o bin/$(PROJECT)-darwin-arm64
	env GOOS=windows GOARCH=amd64 go build -o bin/$(PROJECT)-windows-amd64