// Code generated by ent, DO NOT EDIT.

package review

const (
	// Label holds the string label denoting the review type in the database.
	Label = "review"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// Table holds the table name of the review in the database.
	Table = "reviews"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "reviews"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_reviews"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "reviews"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_reviews"
)

// Columns holds all SQL columns for review fields.
var Columns = []string{
	FieldID,
	FieldBody,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reviews"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_reviews",
	"user_reviews",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// BodyValidator is a validator for the "body" field. It is called by the builders before save.
	BodyValidator func(string) error
)
