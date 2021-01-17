package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
