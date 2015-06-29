package helper

import (
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func init() {

}

func Validate_Json(schema string, document map[string]interface{}) (string, error) {
	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(document)
	validate_result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if validate_result.Valid() {
		fmt.Println("The document is valid\n")
		return "success", nil
	} else {
		fmt.Println("The document is not valid \n")
		for _, desc := range validate_result.Errors() {
			fmt.Printf("- %s\n", desc)
			return "failed", errors.New(desc.Description())
		}
		return "failed", errors.New("unkown errors")
	}
}
