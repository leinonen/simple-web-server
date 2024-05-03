# simple-webserver

A very simple web server built with Go.

## Test your server

```bash
go test
```

## Start your server

```bash
go run main.go
```

## Build the Docker image

```bash
docker build -t simple-webserver .
```

## Run the Docker container

```bash
docker run -p 8080:8080 simple-webserver
```
