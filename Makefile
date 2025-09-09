# Makefile for building, testing, and generating code for the Terraform DataPower provider

# Default target: runs all necessary steps to generate code, documentation, and build binaries
all: generate tfdocs build buildwin

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

# Build the provider binary for the current platform
build:
	go build -o ./.bin/terraform-provider-datapower

# Build the provider binary for Windows (AMD64 architecture)
buildwin:
	GOOS=windows GOARCH=amd64 go build -o ./.bin/terraform-provider-datapower.exe

# Execute code generation templates
execute_templates:
	go run gen/generator.go

# Generate Terraform documentation and format examples
tfdocs:
	go run gen/generate_examples.go
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

# Run acceptance tests with DataPower configuration
test:
	DP_HOSTNAME=datapower-dev DP_USERNAME=admin DP_PASSWORD=admin DP_INSECURE=true DP_ACC_ALL=1 TF_ACC=1 go test -v -p 1 -timeout 90m ./...
