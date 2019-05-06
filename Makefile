.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: install
install: ## format, test, and install main cli binary
	@go generate ./...
	@go fmt ./...
	@go vet ./...
	@go test ./...
	@go install ./...

.PHONY: proto
proto: ## generate protobufs
	@docker run -v `pwd`:/defs namely/prototool:1.17_0 generate

.PHONY: build
build:	## build dockerfile image
	@docker build -t gcr.io/autom8ter-19/tasks .

.PHONY: build-proxy
build-proxy:	## build proxy dockerfile image
	@docker build -t gcr.io/autom8ter-19/tasksproxy -f Dockerfile.proxy .

.PHONY: push-proxy
push-proxy:	## push proxy dockerfile image
	@docker push gcr.io/autom8ter-19/tasksproxy

.PHONY: push
push:	## push docker image
	@docker push gcr.io/autom8ter-19/tasks

.PHONY: run
run:	## run image
	@docker run -d -p 8080:8080 gcr.io/autom8ter-19/tasks


.PHONY: deploy
deploy:	## deploy to kubernetes
	@kubectl apply -f tasks.yaml

.PHONY: deploy-proxy
deploy-proxy:	## deploy to kubernetes(proxy)
	@kubectl apply -f tasksproxy.yaml


.PHONY: update-proxy
update-proxy: ## update kube deployment(proxy)
	@kubectl replace --force -f tasksproxy.yaml

.PHONY: update
update: ## update kube deployment
	@kubectl replace --force -f tasks.yaml
