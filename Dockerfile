
# FROM golang:1.25.2

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/reference/dockerfile/#copy
# COPY *.go ./

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /pipeline-example

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/reference/dockerfile/#expose
# EXPOSE 8080

# # Run
# CMD ["go run main.go"]

FROM golang:1.25.2 AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod .
# COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

# RUN go build -o myapp main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp main.go

# Run the tests in the container
FROM builder AS run-test-stage
RUN go test -v ./...


# FROM alpine:latest as run
FROM alpine:latest 
WORKDIR /app
COPY --from=builder /app/myapp .

# Copy the application executable from the build image
# COPY *.go ./

WORKDIR /app
EXPOSE 8080
CMD ["./myapp"]

