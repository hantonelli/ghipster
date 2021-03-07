package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Text("body").
			NotEmpty(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("reviews").
			Required().
			Unique(),
		edge.From("product", Product.Type).
			Ref("reviews").
			Required().
			Unique(),
	}

}
