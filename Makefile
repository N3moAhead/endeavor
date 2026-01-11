.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o endeavor ./cmd/endeavor/main.go

build-seed: vet
	go build -o seed ./cmd/seed/main.go

run: build
	./endeavor

run-seed: build-seed
	./seed

clean:
	rm -f endeavor
	rm -f seed
