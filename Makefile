GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' Makefile | column -t -s ':' |  sed -e 's/^/ /'

## info: Show project information
info:
	@echo "Git Tag:           ${GIT_TAG}"
	@echo "Git Commit:        ${GIT_COMMIT}"
	@echo "Git Tree State:    ${GIT_DIRTY}"

## generate: Generate code for the project
generate:
	@go generate ./...

## test: Run tests using go test with coverage report
test:
	@"$(CURDIR)/scripts/unit-test.sh"

## staticcheck: Run static check using honnef.co/go/tools/cmd/staticcheck
staticcheck:
	@"$(CURDIR)/scripts/static-check.sh"

## importcheck: Run import check using golang.org/x/tools/cmd/goimports
importcheck:
	@"$(CURDIR)/scripts/import-check.sh"

## fmtcheck: Run format check using go fmt
fmtcheck:
	@"$(CURDIR)/scripts/fmt-check.sh"

## pre-commit: ⚠️ Run all required checks before commit
pre-commit:
	@make staticcheck
	@echo "--------------------------------------------------------------------------------"
	@make importcheck
	@echo "--------------------------------------------------------------------------------"
	@make fmtcheck
	@echo "--------------------------------------------------------------------------------"
	@make test

.NOTPARALLEL:

.PHONY: help