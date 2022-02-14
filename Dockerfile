# syntax=docker/dockerfile:1.3

# Build the Go Binary
FROM registry.suse.com/bci/golang:1.17 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /dc-hello-world

# Create image to bundle app
FROM scratch

COPY --from=build /dc-hello-world /dc-hello-world

EXPOSE 3000

CMD ["/dc-hello-world"]