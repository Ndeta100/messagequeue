
build:
	@echo "Building..."
	@@go build -o bin/messagequeue

clean:
	@echo "Cleaning..."
	@go clean
	@rm -f /bin/messagequeue

test:
	@echo "Testing..."
	@go test -v ./...

run: build
	@echo "Running..."
	@./bin/messagequeue