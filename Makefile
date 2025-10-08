# Makefile for building, testing, and generating code for the Terraform DataPower provider
# 'build' and 'test' commands assume they are being ran in the .devcontainer go environment
# and the datapower container is running with the provided test bed configuration.

# Handle 'make test <subcommand>' where <subcommand> is all (default), resources, or datasources
ifneq ($(filter test,$(MAKECMDGOALS)),)
  TEST_SUBCOMMAND := $(or $(word 2,$(MAKECMDGOALS)),all)
  ifneq ($(word 3,$(MAKECMDGOALS)),)
    $(error Too many arguments for test: $(MAKECMDGOALS))
  endif
  # Swallow the subcommand as a phony target with no action
  $(TEST_SUBCOMMAND):
endif

# Handle 'make build <subcommand>' where <subcommand> is all (default), linux, or windows
ifneq ($(filter build,$(MAKECMDGOALS)),)
  BUILD_SUBCOMMAND := $(or $(word 2,$(MAKECMDGOALS)),all)
  ifneq ($(word 3,$(MAKECMDGOALS)),)
    $(error Too many arguments for build: $(MAKECMDGOALS))
  endif
  # Swallow the subcommand as a phony target with no action
  $(BUILD_SUBCOMMAND):
endif

# Common test environment variables
TEST_ENV = DP_HOSTNAME=datapower-dev DP_USERNAME=admin DP_PASSWORD=admin DP_INSECURE=true DP_ACC_ALL=1 TF_ACC=1

# Default target: runs all necessary steps to generate code, documentation, and build binaries
all: generate tfdocs _build-all

# Clean up generated files and directories
clean:
	bash ./scripts/clean_files.sh ./internal/provider ./gen/generator.ignore

# Generate code, adjust imports, and tidy Go modules
generate: clean execute_templates adjust_imports tidy

# Run goimports to format and fix imports in the provider code
adjust_imports:
	go run golang.org/x/tools/cmd/goimports -w internal/provider/

# Tidy Go module dependencies
tidy:
	go mod tidy

# Execute code generation templates
execute_templates:
	go run gen/generator.go

# Generate Terraform documentation and format examples
tfdocs:
	go run gen/generate_examples.go
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

# Run acceptance tests based on subcommand (all, resources, datasources)
test:
	@$(MAKE) _test-$(TEST_SUBCOMMAND)

# Internal target: Run all acceptance tests
_test-all:
	$(TEST_ENV) go test -v -p 1 -timeout 90m -count=1 ./...

# Internal target: Run tests only for resources
_test-resources:
	$(TEST_ENV) go test -v -p 1 -timeout 90m -count=1 ./internal/provider/resources

# Internal target: Run tests only for datasources
_test-datasources:
	$(TEST_ENV) go test -v -p 1 -timeout 90m -count=1 ./internal/provider/datasources

# Build binaries based on subcommand (all, linux, windows)
build:
	@$(MAKE) _build-$(BUILD_SUBCOMMAND)

# Internal target: Build both Linux and Windows binaries
_build-all: _build-linux _build-windows

# Internal target: Build the provider binary for the current platform (Linux)
_build-linux:
	go build -o ./.bin/terraform-provider-datapower

# Internal target: Build the provider binary for Windows (AMD64 architecture)
_build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./.bin/terraform-provider-datapower.exe

# Declare phony targets to avoid conflicts with files of the same name
.PHONY: all clean generate adjust_imports tidy execute_templates tfdocs test _test-all _test-resources _test-datasources build _build-all _build-linux _build-windows
