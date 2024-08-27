SHELL := /bin/bash
BINARY_NAME=nutanixclient
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build

all: help

## Build:
build: ## Build your project and put the output binary in bin/
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

generate-v3-models: ## Generate V3 models using go-swagger
	swagger generate model \
		--spec=v3/swagger.json \
		--target=v3 \
		--skip-validation \
		--model=prism_central \
		--model=pc_vm \
		--model=mcm_config \
		--model=cmsp_config \
		--model=cmsp_network_config \
		--model=deployment_settings \
		--model=my_ntnx_token \
		--model=cluster_reference \
		--model=pc_vm_nic_configuration \
		--model=network_config \
		--model=subnet \
		--model=availability_zone_reference \
		--model=subnet_resources \
		--model=ip_config \
		--model=network_function_chain_reference \
		--model=virtual_network_reference \
		--model=vpc_reference \
		--model=dhcp_options \
		--model=address \
		--model=ip_pool \
		--model=subnet_def_status \
		--model=subnet_resources_def_status \
		--model=ip_usage_stats \
		--model=pool_stats \
		--model=cluster \
		--model=cluster_network \
		--model=vswitch_config \
		--model=cluster_domain_server \
		--model=credentials \
		--model=cluster_network_entity \
		--model=http_proxy_whitelist \
		--model=smtp_server \
		--model=cluster_config_spec \
		--model=public_key \
		--model=certification_signing_info \
		--model=client_auth \
		--model=encryption_status \
		--model=external_configurations_spec \
		--model=citrix_connector_config_details_spec \
		--model=vm_reference \
		--model=citrix_resource_location_spec \
		--model=cluster_operation_mode \
		--model=cluster_def_status \
		--model=message_resource \
		--model=cluster_analysis \
		--model=cluster_config \
		--model=build_info \
		--model=ca_cert \
		--model=external_configurations \
		--model=cluster_management_server \
		--model=cluster_service_list \
		--model=ssl_key \
		--model=ssl_key_type \
		--model=citrix_connector_config_details \
		--model=citrix_resource_location \
		--model=cluster_nodes \
		--model=hypervisor_server \
		--model=cluster_software \
		--model=software_type

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

