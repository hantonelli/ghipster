// Code generated by entc, DO NOT EDIT.

package gen

import (
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/product"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[0].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
}
