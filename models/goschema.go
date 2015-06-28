package models

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"reflect"
)

func init() {

}

type ApiSchema struct {
	input string `
		{
			"properties": {
				"name": {
					"type": "number"
				}
			},
			"required": ["name"],
			"type": "object"
		}`
	output string `
		{
			"properties": {
				"name": {
					"type": "number"
				}
			},
			"required": ["name"],
			"type": "object"
		}`
}

func TestTag() {
	schema := &ApiSchema{"input_text", "output_text"}

	field, ok := reflect.TypeOf(schema).Elem().FieldByName("input")
	if !ok {
		panic("field not found")
	}
	fmt.Println(string(field.Tag))
	TestGoSchema(string(field.Tag))
}

func TestGoSchema(input string) {
	// schemaLoader := gojsonschema.NewStringLoader(`
	// 	{
	// 		"properties": {
	// 			"name": {
	// 				"type": "number"
	// 			}
	// 		},
	// 		"required": ["name"],
	// 		"type": "object"
	// 	}`)
	schemaLoader := gojsonschema.NewStringLoader(input)
	documentLoader := gojsonschema.NewStringLoader(`{"name" : 3}`)

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

// func ValidInput(input map[string]interface{}{}) {
// 	schemaLoader := gojsonschema.NewStringLoader(`
// 		{
// 			"properties": {
// 				"name": {
// 					"type": "number"
// 				}
// 			},
// 			"required": ["name"],
// 			"type": "object"
// 		}`)
// 	documentLoader := gojsonschema.NewGoLoader(input)

// 	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	if result.Valid() {
// 		fmt.Println("The document is valid\n")
// 	} else {
// 		fmt.Println("The document is not valid \n")
// 		for _, desc := range result.Errors() {
// 			fmt.Printf("- %s\n", desc)
// 		}
// 	}
// }
