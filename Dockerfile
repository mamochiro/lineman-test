# Start from golang base image
FROM golang:1.14.4-alpine as dependencies

ENV GO11MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git make 

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN make pbgen

# Build the Go app
RUN make build
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server cmd/**/*.go

# Start a new stage from scratch
# FROM scratch
FROM alpine

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY templates /app/templates
COPY --from=dependencies /app/bin/server /app/bin/server
COPY --from=dependencies /app/entrypoint.sh /

# # Expose port 8080 to the outside world
# EXPOSE 8080

# #Command to run the executable
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
