FROM golang:1.20-alpine AS builder

LABEL maintainer="Vic Shóstak <vic@shostak.dev> (https://shostak.dev/)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -C cmd/ -ldflags="-s -w" -o apiserver .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/cmd/apiserver", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]