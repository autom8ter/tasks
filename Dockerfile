FROM golang:alpine AS builder

COPY . /go/src/github.com/autom8ter/tasks

RUN set -ex && \
  cd /go/src/github.com/autom8ter/tasks/tasks && \
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./tasks /usr/bin/tasks

FROM busybox
# Retrieve the binary from the previous stage
COPY --from=builder /go/src/github.com/autom8ter/tasks/.env /app/.env
COPY --from=builder /usr/bin/tasks /usr/local/bin/tasks
WORKDIR /app
# Set the binary as the entrypoint of the container
ENTRYPOINT [ "tasks", "serve" ]