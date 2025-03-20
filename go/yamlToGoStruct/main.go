package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var schemas map[string]interface{}
var visited = map[string]bool{}
var output strings.Builder

func main() {
	yamlFile, err := ioutil.ReadFile("input.yaml")
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
		panic("Missing 'components'")
	}

	schemaMap, ok := components["schemas"].(map[string]interface{})
	if !ok {
		panic("Missing 'components.schemas'")
	}
	schemas = schemaMap

	output.WriteString("package main\n\n")

	// Generate structs for all schemas
	for schemaName := range schemas {
		if !visited[schemaName] {
			parseSchema(schemaName)
		}
	}

	err = os.WriteFile("models.go", []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ models.go created successfully.")
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
		for fieldName, fieldValue := range props {
			prop := fieldValue.(map[string]interface{})
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

	if t, ok := prop["type"]; ok {
		switch t {
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
		}
	}

	return "interface{}"
}

func getRefName(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func toCamelCase(input string) string {
	words := strings.Split(input, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
