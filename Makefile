all: build-arm update
.PHONY: all

.PHONY: build-arm
build-arm:
	@echo ">> building arm binary using docker"
	docker run --rm -v $(PWD):/usr/src/go-bind9masq \
      -w /usr/src/go-bind9masq \
      -e GOOS=linux \
      -e GOARCH=arm \
      -e GOARM=7 \
      golang:1.14 go get -d -v ./... && go install -v ./... && go build -o build/go-bind9masq -v

.PHONY: build-amd64
build-amd64:
	@echo ">> building arm binary using docker"
	docker run --rm -v $(PWD):/usr/src/go-bind9masq \
      -w /usr/src/go-bind9masq \
      -e GOOS=linux \
      -e GOARCH=amd64 \
      golang:1.14 go get -d -v ./... && go install -v ./... && go build -o build/go-bind9masq -v

.PHONY: update
update:
	rsync -arpogvt --no-l rsync://ftp.ut-capitole.fr/blacklist .

.PHONY: install
install:
	mv build/go-bind9masq /usr/local/bin/

.PHONY: uninstall
uninstall:
	rm /usr/local/bin/go-bind9masq