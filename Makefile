all: generate tfdocs build buildwin
clean:
	bash ./scripts/clean_files.sh ./internal/provider ./gen/generator.ignore
generate: clean execute_templates adjust_imports tidy
adjust_imports:
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
tidy:
	go mod tidy
build:
	go build -o ./.bin/terraform-provider-datapower
buildwin:
	GOOS=windows GOARCH=amd64 go build -o ./.bin/terraform-provider-datapower.exe
execute_templates:
	go run gen/generator.go
tfdocs:
	go run gen/generate_examples.go
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
test:
	go test -v -p 1 -timeout 90m ./...
