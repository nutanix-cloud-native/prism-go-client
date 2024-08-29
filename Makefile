SHELL := /bin/bash
BINARY_NAME=nutanixclient
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: help

## Build:
.PHONY: build
build: ## Build your project
	mkdir -p bin
	go build ./...

# CRD_OPTIONS define options to add to the CONTROLLER_GEN
CRD_OPTIONS ?= "crd:crdVersions=v1"

.PHONY: run-keploy
run-keploy:
	keploy-server run &

.PHONY: stop-keploy
stop-keploy:
	@-pkill "keploy-server"

generate: $(CONTROLLER_GEN)  ## Generate zz_generated.deepcopy.go
	controller-gen paths="./..." object:headerFile="hack/boilerplate.go.txt"

ORIGINAL_SPEC = v3/openapi/source/original_openapi.json
FIXED_SPEC = v3/openapi/openapi.json
V3_OPENAPI_TEMP_REPO = v3/openapi/temp_repo
PATCHES = $(wildcard v3/openapi/patches/*.patch)

.PHONY: generate-v3-openapi-spec
generate-v3-openapi-spec: $(ORIGINAL_SPEC) $(PATCHES) ## Generate V3 OpenAPI spec by applying bug patches to the original spec
	@echo "Creating temporary Git repository..."
	rm -rf $(V3_OPENAPI_TEMP_REPO)
	mkdir $(V3_OPENAPI_TEMP_REPO)
	cp $(ORIGINAL_SPEC) $(V3_OPENAPI_TEMP_REPO)/
	cd $(V3_OPENAPI_TEMP_REPO) && git init >/dev/null
	cd $(V3_OPENAPI_TEMP_REPO) && git add original_openapi.json
	cd $(V3_OPENAPI_TEMP_REPO) && git commit -m "Original OpenAPI spec" >/dev/null
	@echo "Applying patches..."
	cd $(V3_OPENAPI_TEMP_REPO) && $(foreach patch,$(PATCHES), git apply ../../../$(patch);)
	cd $(V3_OPENAPI_TEMP_REPO) && git add original_openapi.json
	cd $(V3_OPENAPI_TEMP_REPO) && git commit -m "Applied all patches" >/dev/null
	cp $(V3_OPENAPI_TEMP_REPO)/original_openapi.json $(FIXED_SPEC)
	@echo "Fixed OpenAPI spec generated as $(FIXED_SPEC)"
	@echo "Cleaning up..."
	rm -rf $(V3_OPENAPI_TEMP_REPO)
	@echo "Cleaned up"

.PHONY: generate-v3-models
generate-v3-models: ## Generate V3 models using openapi-generator-cli
	openapi-generator-cli generate \
		--input-spec v3/openapi/openapi.json \
		--output v3/models \
		--generator-name go \
		--package-name models \
		--name-mappings uuid=UUID \
		--global-property models \
		--global-property supportingFiles=utils.go \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false

clean: ## Remove build related file
	rm -fr ./bin vendor hack/tools/bin
	rm -f checkstyle-report.xml ./coverage.out ./profile.cov yamllint-checkstyle.xml

## Test:
test: run-keploy ## Run the tests of the project
	go test -race -v ./...
	@$(MAKE) stop-keploy

coverage: run-keploy ## Run the tests of the project and export the coverage
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	@$(MAKE) stop-keploy

## Lint:
lint: lint-go lint-yaml lint-kubebuilder ## Run all available linters

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	golangci-lint run -v

lint-yaml: ## Use yamllint on the yaml file of your projects
ifeq ($(EXPORT_RESULT), true)
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | yamllint-checkstyle > yamllint-checkstyle.xml)
endif
	yamllint -c .yamllint --no-warnings -f parsable $(shell git ls-files '*.yml' '*.yaml') $(OUTPUT_OPTIONS)

.PHONY: lint-kubebuilder
lint-kubebuilder: ## Lint Kubebuilder annotations by generating objects and checking if it is successful
	controller-gen $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=.

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

