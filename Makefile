BASEDIR=${CURDIR}
GOBIN=${BASEDIR}/bin

SERVICE_NAME=hello-berlinski
LOCAL_SERVICE_NAME=${GOBIN}/${SERVICE_NAME}

GO?=go
GOFLAGS=
LDFLAGS=

run:
	$(GO) run cmd/hello-berlinski/main.go

build:
	$(info # Building...)
	@cd cmd/hello-berlinski && $(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS) -w -s" -o $(LOCAL_SERVICE_NAME)

snap:
	snapcraft

