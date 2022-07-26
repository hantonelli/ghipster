package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	// "entgo.io/ent/schema/edge"
	// "entgo.io/ent/schema/index"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Text("name").
			NotEmpty(),
	}
}
