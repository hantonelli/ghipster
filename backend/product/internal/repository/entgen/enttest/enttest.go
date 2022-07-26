// Code generated by ent, DO NOT EDIT.

package enttest

import (
	"context"

	"github.com/hantonelli/ghipster/product/internal/repository/entgen"
	// required by schema hooks.
	_ "github.com/hantonelli/ghipster/product/internal/repository/entgen/runtime"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/hantonelli/ghipster/product/internal/repository/entgen/migrate"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...interface{})
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts        []entgen.Option
		migrateOpts []schema.MigrateOption
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...entgen.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// WithMigrateOptions forwards options to auto migration.
func WithMigrateOptions(opts ...schema.MigrateOption) Option {
	return func(o *options) {
		o.migrateOpts = append(o.migrateOpts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls entgen.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *entgen.Client {
	o := newOptions(opts)
	c, err := entgen.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	migrateSchema(t, c, o)
	return c
}

// NewClient calls entgen.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *entgen.Client {
	o := newOptions(opts)
	c := entgen.NewClient(o.opts...)
	migrateSchema(t, c, o)
	return c
}
func migrateSchema(t TestingT, c *entgen.Client, o *options) {
	tables, err := schema.CopyTables(migrate.Tables)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := migrate.Create(context.Background(), c.Schema, tables, o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
