# Initial stage: download modules
FROM golang:1.15 as modules

RUN mkdir /root/.ssh
ADD go.mod go.sum /m/
RUN cd /m && go mod download

# Intermediate stage: Build the binary
FROM golang:1.15 as builder

COPY --from=modules /go/pkg /go/pkg

# add a non-privileged user
RUN useradd -u 10001 application

RUN mkdir -p /application
ADD . /application
WORKDIR /application

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/app ./cmd/todo/main.go

# Final stage: Run the binary
FROM scratch

# don't forget /etc/passwd from previous stage
COPY --from=builder /etc/passwd /etc/passwd
USER application

# and finally the binary
COPY --from=builder /application/bin/app /app
CMD ["/app"]