IMAGE_NAME = fetcher

CONTAINER_NAME = fetcher

build:
	docker build -t $(IMAGE_NAME) .

# Run with ARGS="--metadata https://www.google.com"
run:
	docker run $(IMAGE_NAME) $(ARGS)

test:
	go test github.com/Pensk/web-fetch/...

.PHONY: build run