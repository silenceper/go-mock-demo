COMMIT_ID = $(shell  git rev-parse --short HEAD)
VERSION ?= $(shell git describe --tags --match='v*' --dirty='.dirty')
REGISTRY = "docker.io"
REPO ?= $(REGISTRY)/silenceper
TAG ?= $(VERSION)-$(COMMIT_ID)

default: image
GOOS ?= linux
GOARCH ?= amd64

# Run go fmt against code
.PHONY: fmt
fmt:
	@find . -type f -name '*.go'| grep -v "/vendor/" | xargs gofmt -w -s

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...
# Run mod tidy against code
.PHONY: tidy
tidy:
	@go mod tidy

# Run go mod vendor
.PHONY: vendor
vendor:
	@go mod vendor
.PHONY: build
build:
	go build -o bin/go-mock-demo main.go


.PHONY: push
push: tidy fmt
	docker buildx build --platform linux/amd64 --build-arg VERSION=$(VERSION) -f ./Dockerfile -t $(REPO)/go-mock-demo:$(TAG) --push .