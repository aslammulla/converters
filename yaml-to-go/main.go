package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	schemas map[string]interface{}
	visited = map[string]bool{}
	output  strings.Builder
)

func main() {
	// Read YAML file
	yamlFile, err := ioutil.ReadFile("openapi.yaml")
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	components, ok := data["components"].(map[string]interface{})
	if !ok {
		panic("Missing components")
	}

	schemaMap, ok := components["schemas"].(map[string]interface{})
	if !ok {
		panic("Missing components.schemas")
	}
	schemas = schemaMap

	output.WriteString("package main\n\n")

	for name := range schemas {
		if !visited[name] {
			parseSchema(name)
		}
	}

	err = os.WriteFile("models.go", []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ models.go generated successfully.")
}

func parseSchema(name string) {
	if visited[name] {
		return
	}
	visited[name] = true

	schemaData := schemas[name].(map[string]interface{})
	props, hasProps := schemaData["properties"].(map[string]interface{})

	output.WriteString(fmt.Sprintf("type %s struct {\n", name))

	if hasProps {
		for fieldName, fieldVal := range props {
			prop := fieldVal.(map[string]interface{})
			goType := resolveType(prop)
			output.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", toCamelCase(fieldName), goType, fieldName))
		}
	}

	output.WriteString("}\n\n")
}

func resolveType(prop map[string]interface{}) string {
	if ref, ok := prop["$ref"]; ok {
		refName := getRefName(ref.(string))
		parseSchema(refName)
		return refName
	}

	switch prop["type"] {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	case "number":
		return "float64"
	case "array":
		if items, ok := prop["items"].(map[string]interface{}); ok {
			return "[]" + resolveType(items)
		}
		return "[]interface{}"
	case "object":
		if additionalProps, ok := prop["additionalProperties"].(map[string]interface{}); ok {
			return "map[string]" + resolveType(additionalProps)
		}
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

func getRefName(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func toCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
