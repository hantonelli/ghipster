// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/entgen/product"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/entgen/review"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/entgen/user"
)

// ReviewCreate is the builder for creating a Review entity.
type ReviewCreate struct {
	config
	mutation *ReviewMutation
	hooks    []Hook
}

// SetBody sets the "body" field.
func (rc *ReviewCreate) SetBody(s string) *ReviewCreate {
	rc.mutation.SetBody(s)
	return rc
}

// SetID sets the "id" field.
func (rc *ReviewCreate) SetID(i int) *ReviewCreate {
	rc.mutation.SetID(i)
	return rc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (rc *ReviewCreate) SetAuthorID(id int) *ReviewCreate {
	rc.mutation.SetAuthorID(id)
	return rc
}

// SetAuthor sets the "author" edge to the User entity.
func (rc *ReviewCreate) SetAuthor(u *User) *ReviewCreate {
	return rc.SetAuthorID(u.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (rc *ReviewCreate) SetProductID(id int) *ReviewCreate {
	rc.mutation.SetProductID(id)
	return rc
}

// SetProduct sets the "product" edge to the Product entity.
func (rc *ReviewCreate) SetProduct(p *Product) *ReviewCreate {
	return rc.SetProductID(p.ID)
}

// Mutation returns the ReviewMutation object of the builder.
func (rc *ReviewCreate) Mutation() *ReviewMutation {
	return rc.mutation
}

// Save creates the Review in the database.
func (rc *ReviewCreate) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("entgen: uninitialized hook (forgotten import entgen/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Review)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReviewMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReviewCreate) SaveX(ctx context.Context) *Review {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReviewCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReviewCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReviewCreate) check() error {
	if _, ok := rc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`entgen: missing required field "Review.body"`)}
	}
	if v, ok := rc.mutation.Body(); ok {
		if err := review.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`entgen: validator failed for field "Review.body": %w`, err)}
		}
	}
	if _, ok := rc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`entgen: missing required edge "Review.author"`)}
	}
	if _, ok := rc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product", err: errors.New(`entgen: missing required edge "Review.product"`)}
	}
	return nil
}

func (rc *ReviewCreate) sqlSave(ctx context.Context) (*Review, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (rc *ReviewCreate) createSpec() (*Review, *sqlgraph.CreateSpec) {
	var (
		_node = &Review{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: review.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldBody,
		})
		_node.Body = value
	}
	if nodes := rc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.AuthorTable,
			Columns: []string{review.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_reviews = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.ProductTable,
			Columns: []string{review.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_reviews = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReviewCreateBulk is the builder for creating many Review entities in bulk.
type ReviewCreateBulk struct {
	config
	builders []*ReviewCreate
}

// Save creates the Review entities in the database.
func (rcb *ReviewCreateBulk) Save(ctx context.Context) ([]*Review, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Review, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReviewMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReviewCreateBulk) SaveX(ctx context.Context) []*Review {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReviewCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReviewCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}