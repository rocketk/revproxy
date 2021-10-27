PROJECT=revproxy

.PHONY: build

build:
	# for the current OS only
	go build
buildwin:
	# for windows-amd64 only
	env GOOS=windows GOARCH=amd64 go build -o bin/$(PROJECT)-linux-amd64
buildlinux: 
	# for linux-amd64 only
	env GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT)-linux-amd64
