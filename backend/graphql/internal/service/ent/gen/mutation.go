// Code generated by entc, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"sync"

	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/predicate"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/product"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProduct = "Product"
)

// ProductMutation represents an operation that mutates the Product nodes in the graph.
type ProductMutation struct {
	config
	op            Op
	typ           string
	id            *int
	text          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Product, error)
	predicates    []predicate.Product
}

var _ ent.Mutation = (*ProductMutation)(nil)

// productOption allows management of the mutation configuration using functional options.
type productOption func(*ProductMutation)

// newProductMutation creates new mutation for the Product entity.
func newProductMutation(c config, op Op, opts ...productOption) *ProductMutation {
	m := &ProductMutation{
		config:        c,
		op:            op,
		typ:           TypeProduct,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProductID sets the ID field of the mutation.
func withProductID(id int) productOption {
	return func(m *ProductMutation) {
		var (
			err   error
			once  sync.Once
			value *Product
		)
		m.oldValue = func(ctx context.Context) (*Product, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Product.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProduct sets the old Product of the mutation.
func withProduct(node *Product) productOption {
	return func(m *ProductMutation) {
		m.oldValue = func(context.Context) (*Product, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProductMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProductMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("gen: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *ProductMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetText sets the "text" field.
func (m *ProductMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *ProductMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Product entity.
// If the Product object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProductMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *ProductMutation) ResetText() {
	m.text = nil
}

// Op returns the operation name.
func (m *ProductMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Product).
func (m *ProductMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProductMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.text != nil {
		fields = append(fields, product.FieldText)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProductMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case product.FieldText:
		return m.Text()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProductMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case product.FieldText:
		return m.OldText(ctx)
	}
	return nil, fmt.Errorf("unknown Product field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) SetField(name string, value ent.Value) error {
	switch name {
	case product.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProductMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProductMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Product numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProductMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProductMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProductMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Product nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProductMutation) ResetField(name string) error {
	switch name {
	case product.FieldText:
		m.ResetText()
		return nil
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProductMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProductMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProductMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProductMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProductMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProductMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProductMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Product unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProductMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Product edge %s", name)
}
