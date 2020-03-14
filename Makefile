REVIEWDOG_ARG ?= -diff="git diff master"
NAME = terraform-provider-cmdexec

LINT_TOOLS=\
	github.com/reviewdog/reviewdog/cmd/reviewdog \
	golang.org/x/lint/golint \
	honnef.co/go/tools/cmd/staticcheck \
	github.com/quasilyte/go-consistent \
 	github.com/securego/gosec/cmd/gosec

.PHONY: all
all: test reviewdog

.PHONY: bootstrap-lint-tools
bootstrap-lint-tools:
	@for tool in $(LINT_TOOLS) ; do \
		echo "Installing $$tool" ; \
		go get -u $$tool; \
	done

.PHONY: fmtcheck
fmtcheck:
	@go fmt ./... | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

.PHONY: mod
mod:
	go mod download

.PHONY: test
test: fmtcheck
	TF_ACC=1 go test ./... -v
	go test -v -race ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: reviewdog
reviewdog:
	reviewdog -conf=.reviewdog.yml $(REVIEWDOG_ARG)

.PHONY: build
build: fmtcheck vet
	go build -o $(NAME) .

.PHONY: build-linux
build-linux: fmtcheck vet
	GOOS=linux GOARCH=amd64 go build -o $(NAME) .

.PHONY: install
install: build
	mv $(NAME) ~/.terraform.d/plugins/
