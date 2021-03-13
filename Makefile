all: build-amd64
.PHONY: all

BINARY = "go-bind9masq"

.PHONY: build-arm
build-arm:
	@echo ">> building arm binary using docker"
	docker run --rm -v $(PWD):/usr/src/$(BINARY) \
      -w /usr/src/$(BINARY) \
      -e GOOS=linux \
      -e GOARCH=arm \
      -e GOARM=7 \
      golang:1.14 go get -d -v ./... && go install -v ./... && go build -v -o build/$(BINARY) -v ./...

.PHONY: build-amd64
build-amd64:
	@echo ">> building amd64 binary using docker"
	docker run --rm -v $(PWD):/usr/src/$(BINARY) \
      -w /usr/src/$(BINARY) \
      -e GOOS=linux \
      -e GOARCH=amd64 \
      golang:1.14 go get -d -v ./... && go install -v ./... && go build -v -o build/$(BINARY) -v ./...

.PHONY: install
install:
	chmod 755 build/$(BINARY)
	sudo cp build/$(BINARY) /usr/local/bin/
	sudo mkdir /etc/$(BINARY)
	sudo cp etc/config.yml /etc/$(BINARY)/config.yml
	sudo cp etc/$(BINARY)-update /etc/cron.daily/

.PHONY: uninstall
uninstall:
	sudo rm /usr/local/bin/$(BINARY)
	sudo rm /etc/cron.daily/$(BINARY)-update
