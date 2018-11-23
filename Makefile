

.PHONY: clean build-goroutine-example-1

clean:
	go clean
	rm -rf build/

build-goroutine-example-1:
	@echo build-goroutine-example1
	go build -o build/goroutine-example-1 cmd/goroutine-example-1/main.go
