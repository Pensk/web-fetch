# Web Fetcher

CLI for saving html and parsing metadata on those files.

## Prerequisites

Docker and Make or Go

## Makefile Rules

- `build`: Builds the Docker image.
- `run`: Runs the Docker container. 

## Build and Run the Project

### Using Docker

1. Build the Docker image:

    ```bash
    make build
    ```

2. Run the Docker container:

    ```bash
    make run ARGS="--metadata https://www.google.com"
    ```

### Using Go

1. Build the application:

    ```bash
    go build -o fetch cmd/main.go
    ```

2. Run the application:

    ```bash
    ./fetch --metadata https://www.google.com
    ```

### Test

```bash
make test
```