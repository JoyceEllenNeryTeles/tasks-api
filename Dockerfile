# syntax=docker/dockerfile:1

FROM golang:1.24.3

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 go get -a -ldflags '-s' github.com/JoyceEllenNeryTeles/test/tasks-api

COPY --from=builder /go/bin/tasks-api .

EXPOSE 8080

# Run
CMD ["/tasks-api"]