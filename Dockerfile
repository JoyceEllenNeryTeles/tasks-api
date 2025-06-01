# syntax=docker/dockerfile:1

FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 go get -a -ldflags '-s' github.com/JoyceEllenNeryTeles/test/tasks-api

COPY --from=builder /go/bin/tasks-api .

EXPOSE 8080

# Run
CMD ["/tasks-api"]