📦 Go Struct Generator

This repository contains two utilities to generate Go structs from input files:

    ✅ JSON to Go Struct: Convert a raw JSON file into Go struct definitions.
    ✅ YAML to Go Struct (OpenAPI 3.0): Convert OpenAPI YAML (with $ref schemas) into nested Go structs.

📁 Folder Structure

.
├── json-to-go/
│   ├── input.json
│   └── main.go
│
├── yaml-to-go/
│   ├── input.yaml
│   └── main.go
│
└── models.go  # Output file (generated)

🔧 Requirements

    Go 1.18+
    gopkg.in/yaml.v3 (for YAML parsing)

Install the YAML package:

go get gopkg.in/yaml.v3

✅ JSON to Go Struct
🔹 Description

Parses a JSON file and converts it into Go struct definitions. Handles nested objects and arrays.
🔹 Usage

    Place your input JSON in json-to-go/input.json
    Run the converter:

cd json-to-go
go run main.go

    The result will be printed to console or written to models.go

✅ YAML (OpenAPI) to Go Struct
🔹 Description

Reads an OpenAPI 3.0 YAML file and converts all schemas (in components.schemas) into separate Go structs.

    Supports $ref: '#/components/schemas/XYZ'
    Handles nested and recursive references
    Supports object, string, int, float, array, map types

🔹 Usage

    Place your OpenAPI file in yaml-to-go/input.yaml
    Run the converter:

cd yaml-to-go
go run main.go

    The result will be saved in models.go

📌 Output Example

type Customer struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
}