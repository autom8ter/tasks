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
- [x] GRPC API (Kubernetes)
- [x] Rest Gateway API (Kubernetes)

### Documentation
- [x] [GRPC Protocol Documentation]()

### Build/CI/Deployment

- [x] [GRPC API Dockerfile](https://github.com/autom8ter/tasks/blob/master/Dockerfile)
- [x] [GRPC API Kubernetes Manifest](https://github.com/autom8ter/tasks/blob/master/tasks.yaml)
- [x] [Rest Gateway API Dockerfile](https://github.com/autom8ter/tasks/blob/master/Dockerfile.proxy)
- [x] [Rest Gateway API Kubernetes Manifest](https://github.com/autom8ter/tasks/blob/master/tasksproxy.yaml)
- [x] [Makefile](https://github.com/autom8ter/tasks/blob/master/Makefile)
- [x] [Prototool-Protobuf build tool](https://github.com/autom8ter/tasks/blob/master/prototool.yaml)

### SDK's
- [x] [C++ SDK](https://github.com/autom8ter/tasks/tree/master/sdk/cpp/tasks)
- [x] [C# SDK](https://github.com/autom8ter/tasks/tree/master/sdk/csharp/tasks)
- [x] [Go SDK](https://github.com/autom8ter/tasks/tree/master/sdk/go/tasks)
- [x] [Java SDK](https://github.com/autom8ter/tasks/tree/master/sdk/java/tasks/tasks)
- [x] [Javascript SDK](https://github.com/autom8ter/tasks/tree/master/sdk/js/tasks)
- [x] [Objective-C SDK](https://github.com/autom8ter/tasks/tree/master/sdk/objc/tasks)
- [x] [Python SDK](https://github.com/autom8ter/tasks/tree/master/sdk/python/tasks)
- [x] [Ruby SDK](https://github.com/autom8ter/tasks/tree/master/sdk/ruby/tasks)

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

