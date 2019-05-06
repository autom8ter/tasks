# Tasks

    go get github.com/autom8ter/tasks/...

**Documentation is a WIP, I am on vacation 5/3-5/8**

- gRPC API address(for client sdk's): 104.198.16.50:8080

- REST API address(this is a reverse proxy for the gRPC API for web clients): 35.238.49.10:8080

## REST endpoints:
- http://35.238.49.10:8080/create
- http://35.238.49.10:8080/read
- http://35.238.49.10:8080/update
- http://35.238.49.10:8080/delete
- http://35.238.49.10:8080/list

- http://35.238.49.10:8080/docs
- http://35.238.49.10:8080/metrics
- http://35.238.49.10:8080/debug/pprof/



## Features

### Independent Services
- [x] Postgres (Kubernetes)
- [x] GRPC API (Kubernetes) endpoint: 
- [x] Rest Gateway API (Kubernetes) endpoint:

### Documentation
- [x] [GRPC Protocol Documentation]()

### Build/CI/Deployment

- [x] [GRPC API Dockerfile]()
- [x] [GRPC API Kubernetes Manifest]()
- [x] [Rest Gateway API Dockerfile]()
- [x] [Rest Gateway API Kubernetes Manifest]()
- [x] [Makefile]()
- [x] [Prototool-Protobuf build tool]()

### SDK's
- [x] [C++ SDK]()
- [x] [C# SDK]()
- [x] [Go SDK]()
- [x] [Java SDK]()
- [x] [Javascript SDK]()
- [x] [Objective-C SDK]()
- [x] [Python SDK]()
- [x] [Ruby SDK]()

### Methods
- [x] Create Task   
- [x] Read Task    
- [x] Update Task  
- [x] Delete Task   
- [x] List Tasks   

### Metrics
 
- [x] Total Requests
- [x] Request Duration
- [x] In-Flight Requests
- [x] Response Size

### Profiling
- [x] pprof/
- [x] pprof/cmdline
- [x] pprof/profile
- [x] pprof/symbol
- [x] pprof/trace

### Dependency Management
- [x] go modules


## Makefile

run:  `make help` in the root directory of this project

output:

```text
help:    show this help
install:  format, test, and install main cli binary
proto:  generate protobufs
build:   build dockerfile image
build-proxy:     build proxy dockerfile image
push-proxy:      push proxy dockerfile image
push:    push docker image
run:     run image
deploy:  deploy to kubernetes
deploy-proxy:    deploy to kubernetes(proxy)
update-proxy:  update kube deployment(proxy)
update:  update kube deployment

```


## TODO
- [ ] unit tests
- [ ] integration tests
- [ ] documentation for all functions
- [ ] stackdriver logging
- [ ] add more mock data
- [ ] add example curl commands
- [ ] add example grpc client services

