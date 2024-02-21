package openapi

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

func getOperation(pathItem *spec.PathItem, method *string) (*string, error) {
	getFieldValue := reflect.ValueOf(*pathItem).FieldByName(*method)

	fmt.Println("a", getFieldValue)

	if getFieldValue.IsValid() {
		// Get the Operation struct from the field value
		operation := getFieldValue.Interface().(*spec.Operation)
		fmt.Println("operation to perform: ", operation.ID)
		fmt.Println("Description: ", operation.Description)
		return &operation.ID, nil
	} else {
		operation := ""
		return &operation, fmt.Errorf("%s method not found in spec", *method)
	}
}

func LookupSpec(path string, method *string) (*string, error) {
	// Load the OpenAPI specification file
	swaggerSpec, err := loads.Spec("swagger.yaml")
	if err != nil {
		fmt.Println("Failed to load Swagger spec: ", err)
	}

	// Get the paths from the OpenAPI specification
	p := swaggerSpec.Spec().Paths.Paths[path]
	return getOperation(&p, method)
}
