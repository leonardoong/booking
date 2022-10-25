-include .env

.PHONY: build

# Build
build: build-http

build-http:
	@echo " > Building [booking-system]"
	@cd ./cmd/http/ && go build -o ../../bin.exe && cd ../..
	@echo " > Finished building [booking-system]"

# Run
run-http:
	@echo " > Running [booking-system]"
	@cd ./cmd/http/ && go build -o ../../bin.exe && cd ../..
	@./bin.exe
	@echo " > Finished running [booking-system]"

# Test
test:
	@echo " > Testing all..."
	@go test -race -cover ./
	@echo " > Finished Testing"