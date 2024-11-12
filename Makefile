# Default to docker if no ENGINE is specified
ENGINE ?= docker

# Image name
IMAGE_NAME = pid_teller

# Common targets that work for both engines
.PHONY: build run clean help

build:
	$(ENGINE) build -t $(IMAGE_NAME) .

run:
	$(ENGINE) run -d --name $(IMAGE_NAME) $(IMAGE_NAME)

clean:
	$(ENGINE) stop $(IMAGE_NAME) 2>/dev/null || true
	$(ENGINE) rm $(IMAGE_NAME) 2>/dev/null || true

help:
	@echo "Usage:"
	@echo "  make [target] [ENGINE=docker|podman]"
	@echo ""
	@echo "Targets:"
	@echo "  build   - Build the image"
	@echo "  run     - Run the container in detached mode"
	@echo "  clean   - Stop and remove the container"
	@echo ""
	@echo "Examples:"
	@echo "  make build ENGINE=docker"
	@echo "  make run ENGINE=podman"

build:
	go build main.go

run-local:
	go run main.go

run-unshare:
	sudo unshare --pid --fork --mount-proc ./main


build-docker:
	docker build -t nsenter-demo .

run-docker:
	docker run --name nsenter -d nsenter-demo