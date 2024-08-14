# syntax=docker/dockerfile:1

FROM golang:1.21.3

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./

# Build the Go app
RUN go build -o restful-api 

# Run the Go app
CMD ["./restful-api"]