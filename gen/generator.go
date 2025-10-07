// Copyright © 2025 Scott Wiederhold <s.e.wiederhold@gmail.com>
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"gopkg.in/yaml.v3"
)

const (
	modelDefinitionsPath  = "./gen/definitions"
	exGenTemplate         = "./gen/templates/generate_examples.go.tmpl"
	exGenLocation         = "./gen/generate_examples.go"
	providerTemplate      = "./gen/templates/provider.go.tmpl"
	providerLocation      = "./internal/provider/provider.go"
	changelogTemplate     = "./gen/templates/changelog.md.tmpl"
	changelogLocation     = "./templates/guides/changelog.md.tmpl"
	changelogOriginal     = "./CHANGELOG.md"
	actionMapTemplate     = "./gen/templates/actionmap.go.tmpl"
	actionMapLocation     = "./internal/provider/actions/actionmap.go"
	actionMapTestTemplate = "./gen/templates/actionmap_test.go.tmpl"
	actionMapTestLocation = "./internal/provider/actions/actionmap_test.go"
	testConfigTemplate    = "./gen/templates/testconfig.go.tmpl"
	testConfigLocation    = "./testutils/config.go"
)

type templateInfo struct {
	path   string
	prefix string
	suffix string
}

var providerTemplates = []templateInfo{
	{
		path:   "./gen/templates/model.go.tmpl",
		prefix: "./internal/provider/models/model_datapower_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go.tmpl",
		prefix: "./internal/provider/datasources/data_source_datapower_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go.tmpl",
		prefix: "./internal/provider/datasources/data_source_datapower_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go.tmpl",
		prefix: "./internal/provider/resources/resource_datapower_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go.tmpl",
		prefix: "./internal/provider/resources/resource_datapower_",
		suffix: "_test.go",
	},
}

type YamlConfig struct {
	Name              string                `yaml:"name"`
	ModelOnly         bool                  `yaml:"model_only"`
	DefaultFunction   bool                  `yaml:"default_function"`
	ListItem          bool                  `yaml:"list_item"`
	RestEndpoint      string                `yaml:"rest_endpoint"`
	Singleton         bool                  `yaml:"singleton"`
	WriteOnlyAttrs    bool                  `yaml:"write_only_attrs"`
	DefaultDomainOnly bool                  `yaml:"default_domain_only"`
	Description       string                `yaml:"description"`
	Deprecated        string                `yaml:"deprecated"`
	CliAlias          string                `yaml:"cli_alias"`
	SkipTest          bool                  `yaml:"skip_test"`
	TestPrerequisites string                `yaml:"test_prerequisites"`
	Actions           []YamlActionAttribute `yaml:"actions"`
	Attributes        []YamlConfigAttribute `yaml:"attributes"`
	DataResource      []string
}

type YamlValidationTest struct {
	Evaluation  string               `yaml:"evaluation"`
	Conditions  []YamlValidationTest `yaml:"conditions"`
	Attribute   string               `yaml:"attribute"`
	AttrTfName  string               `yaml:"attr_tf_name"`
	AttrPath    string               `yaml:"attr_path"`
	AttrType    string               `yaml:"attr_type"`
	AttrDefault string               `yaml:"attr_default"`
	Value       []string             `yaml:"value"`
}

type YamlActionAttribute struct {
	Name string `yaml:"name"`
	Body string `yaml:"body"`
}

type YamlConfigAttribute struct {
	Name            string               `yaml:"name"`
	Path            string               `yaml:"path"`
	TfName          string               `yaml:"tf_name"`
	Type            string               `yaml:"type"`
	BoolType        string               `yaml:"bool_type"`
	DmType          string               `yaml:"dm_type"`
	ElementType     string               `yaml:"element_type"`
	Internal        bool                 `yaml:"internal"`
	Required        bool                 `yaml:"required"`
	Computed        bool                 `yaml:"computed"`
	WriteOnly       bool                 `yaml:"write_only"`
	WriteOnlyAttrs  bool                 `yaml:"write_only_attrs"`
	Description     string               `yaml:"description"`
	Deprecated      string               `yaml:"deprecated"`
	Sensitive       string               `yaml:"sensitive"`
	CliAlias        string               `yaml:"cli_alias"`
	Example         string               `yaml:"example"`
	ListExample     []string             `yaml:"list_example"`
	Enum            []string             `yaml:"enum"`
	Minimum         int64                `yaml:"minimum"`
	Maximum         int64                `yaml:"maximum"`
	StringPatterns  []string             `yaml:"string_patterns"`
	StringExclude   []string             `yaml:"string_exclude"`
	StringMinLength int64                `yaml:"string_min_length"`
	StringMaxLength int64                `yaml:"string_max_length"`
	Default         string               `yaml:"default"`
	ListDefault     []string             `yaml:"list_default"`
	ResetValue      string               `yaml:"reset_value"`
	TestValue       string               `yaml:"test_value"`
	TestBedValue    string               `yaml:"test_bed_value"`
	TestExpectDs    string               `yaml:"test_expect_ds"`
	ReferenceTo     string               `yaml:"reference_to"`
	FileLocations   []string             `yaml:"file_locations"`
	Protocols       []string             `yaml:"protocols"`
	RequiredWhen    []YamlValidationTest `yaml:"required_when"`
	IgnoredWhen     []YamlValidationTest `yaml:"ignored_when"`
}

// getAttributeDefault returns attribute default value
func getAttributeDefault(attributes []YamlConfigAttribute, name string) string {
	for _, attr := range attributes {
		if attr.Name == name {
			return attr.Default
		}
	}
	return "ERROR-NOT-FOUND"
}

// getAttributeTfName returns attribute TfName
func getAttributeTfName(attributes []YamlConfigAttribute, name string) string {
	for _, attr := range attributes {
		if attr.Name == name {
			return attr.TfName
		}
	}
	return tfutils.ToTfName(name)
}

// getAttributeType returns attribute type or dm_type value
func getAttributeType(attributes []YamlConfigAttribute, name string) string {
	for _, attr := range attributes {
		if attr.Name == name {
			if attr.Type != "" {
				return attr.Type
			} else if attr.DmType != "" {
				return attr.DmType
			}
		}
	}
	return "ERROR-NOT-FOUND"
}

// hasDomainAttribute returns true if an "AppDomain" attribute is present.
func hasDomainAttribute(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Name == "AppDomain" {
			return true
		}
	}
	return false
}

// hasIdAttribute returns true if an "Id" attribute is present.
func hasIdAttribute(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Name == "Id" {
			return true
		}
	}
	return false
}

// isList returns true if the attribute is a list without nested elements.
func isList(attribute YamlConfigAttribute) bool {
	return attribute.Type == "List" && attribute.ElementType != ""
}

// isObject returns true if the attribute is an object (not list, string, or int64).
func isObject(attribute YamlConfigAttribute) bool {
	return attribute.Type != "List" && attribute.Type != "String" && attribute.Type != "Int64" && attribute.Type != "Bool"
}

// isObjectList returns true if the attribute is a list with nested objects.
func isObjectList(attribute YamlConfigAttribute) bool {
	return attribute.Type == "List" && attribute.ElementType != "String"
}

// isStringList returns true if the attribute is a list of strings.
func isStringList(attribute YamlConfigAttribute) bool {
	return attribute.Type == "List" && attribute.ElementType == "String"
}

// toGoName converts a snake_case string to Go-style CamelCase.
func toGoName(s string) string {
	var b strings.Builder
	for part := range strings.SplitSeq(s, "_") {
		b.WriteString(cases.Title(language.English).String(part))
	}
	return b.String()
}

// updateComputed returns true if Computed or WriteOnly are set on any attributes.
func updateComputed(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Computed || attr.WriteOnly {
			return true
		}
	}
	return false
}

// Map of templating functions.
var functions = template.FuncMap{
	"getAttributeDefault": getAttributeDefault,
	"getAttributeTfName":  getAttributeTfName,
	"getAttributeType":    getAttributeType,
	"hasDomainAttribute":  hasDomainAttribute,
	"hasIdAttribute":      hasIdAttribute,
	"isList":              isList,
	"isObject":            isObject,
	"isObjectList":        isObjectList,
	"isStringList":        isStringList,
	"lower":               strings.ToLower,
	"updateComputed":      updateComputed,
	"strContains":         strings.Contains,
	"toGoName":            toGoName,
	"toTfName":            tfutils.ToTfName,
	"quote":               strconv.Quote,
}

// buildAttribute derives TfName from Name if not set, handling camel to snake conversion
func buildAttribute(attr *YamlConfigAttribute) {
	if attr.Path == "" {
		attr.Path = fmt.Sprintf("`%s`", strings.ReplaceAll(strings.ReplaceAll(attr.Name, ":", `\:`), ".", `\.`))
	} else {
		attr.Path = fmt.Sprintf("`%s`", attr.Path)
	}
	if attr.TfName == "" {
		attr.TfName = tfutils.ToTfName(attr.Name)
	}
	if attr.WriteOnly {
		attr.TfName = attr.TfName + "_wo"
	}
}

// buildConfig augments the config and its attributes.
func buildConfig(config *YamlConfig) {
	config.DataResource = []string{"DataSource", "Resource"}
	for i := range config.Attributes {
		buildAttribute(&config.Attributes[i])
	}
}

// processTemplate renders a template to the output path, creating directories if needed.
func processTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// Skip first line with 'build-ignore' directive for go.tmpl files.
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go.tmpl") {
		scanner.Scan()
	}
	var temp strings.Builder
	for scanner.Scan() {
		temp.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading template: %v", err)
	}

	tmpl, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp.String())
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	if err := tmpl.Execute(output, config); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Create directories if needed.
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	// Write to output file.
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer f.Close()
	if _, err := f.Write(output.Bytes()); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}
}

func main() {
	// Read model definitions directory.
	entries, err := os.ReadDir(modelDefinitionsPath)
	if err != nil {
		log.Fatalf("Error reading definitions directory: %v", err)
	}

	var configs []YamlConfig
	for _, entry := range entries {
		if entry.IsDir() || (!strings.HasSuffix(entry.Name(), ".yaml") && !strings.HasSuffix(entry.Name(), ".yml")) {
			continue
		}
		yamlPath := filepath.Join(modelDefinitionsPath, entry.Name())
		yamlData, err := os.ReadFile(yamlPath)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", yamlPath, err)
		}

		var config YamlConfig
		if err := yaml.Unmarshal(yamlData, &config); err != nil {
			log.Fatalf("Error parsing YAML %s: %v", entry.Name(), err)
		}
		configs = append(configs, config)
	}

	// Sort configs by name for deterministic generation.
	sort.Slice(configs, func(i, j int) bool {
		return configs[i].Name < configs[j].Name
	})

	for i := range configs {
		// Build config.
		buildConfig(&configs[i])

		// Process templates for this config.
		for _, t := range providerTemplates {
			if configs[i].ModelOnly && (strings.Contains(t.path, "data_source") || strings.Contains(t.path, "resource") || strings.Contains(t.path, "data-source.tf") || strings.Contains(t.path, "resource.tf")) {
				continue
			}
			outputPath := t.prefix + tfutils.ToTfName(configs[i].Name) + t.suffix
			processTemplate(t.path, outputPath, configs[i])
		}
	}

	// Process provider.go with all configs.
	processTemplate(providerTemplate, providerLocation, configs)

	// Process actions.go with all configs.
	processTemplate(actionMapTemplate, actionMapLocation, configs)

	// Process actions.go with all configs.
	processTemplate(actionMapTestTemplate, actionMapTestLocation, configs)

	// Process testconfig.go with all configs.
	processTemplate(testConfigTemplate, testConfigLocation, configs)

	// Process example generator with all configs.
	processTemplate(exGenTemplate, exGenLocation, configs)

	// Process changelog.
	changelogData, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	processTemplate(changelogTemplate, changelogLocation, string(changelogData))
}
