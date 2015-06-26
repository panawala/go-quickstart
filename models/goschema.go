package models

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func init() {

}

func TestGoSchema() {
	schemaLoader := gojsonschema.NewStringLoader(`
		{
			"properties": {
				"name": {
					"type": "number"
				}
			},
			"required": ["name"],
			"type": "object" 
		}`)
	documentLoader := gojsonschema.NewStringLoader(`{"name" :"hello world"}`)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println(err.Error())
	}

	if result.Valid() {
		fmt.Println("The document is valid\n")
	} else {
		fmt.Println("The document is not valid \n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func ValidInput(input map[string]interface{}{}) {
	schemaLoader := gojsonschema.NewStringLoader(`
		{
			"properties": {
				"name": {
					"type": "number"
				}
			},
			"required": ["name"],
			"type": "object" 
		}`)
	documentLoader := gojsonschema.NewGoLoader(input)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println(err.Error())
	}

	if result.Valid() {
		fmt.Println("The document is valid\n")
	} else {
		fmt.Println("The document is not valid \n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}