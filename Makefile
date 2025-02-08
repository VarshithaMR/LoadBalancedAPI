# Name of binary executable
NAME=unique-id-processor

# Local path
LOCAL_ENV=PATH="$(GOPATH)/bin:$(PATH)"

# Set Go module
GO=GO111MODULE=on $(LOCAL_ENV) go

#Root files
ROOT_FILE=cmd/main.go

#set c-go, OS and Arch
CGO=0
GOOS=linux
GOARCH=amd64

# build the service
build:
	@echo "Building calculate-application-go....."
	$(GO) mod tidy
	CGO_ENABLED=${CGO} GOOS=${GOOS} GOARCH=${GOARCH} $(GO) build -o $(NAME) $(ROOT_FILE)
.PHONY: build

# run the service
run:
	@echo "Running the application...."
	$(GO) run $(ROOT_FILE)
.PHONY: run