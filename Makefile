.PHONY: fronend
fronend:
	cd fronend && yarn dev

build:
	go generate ./...
	go build -ldflags="-X main.BuildEnv=prod"
