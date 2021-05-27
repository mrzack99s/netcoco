// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/devicetype"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// DeviceTypeDelete is the builder for deleting a DeviceType entity.
type DeviceTypeDelete struct {
	config
	hooks    []Hook
	mutation *DeviceTypeMutation
}

// Where adds a new predicate to the DeviceTypeDelete builder.
func (dtd *DeviceTypeDelete) Where(ps ...predicate.DeviceType) *DeviceTypeDelete {
	dtd.mutation.predicates = append(dtd.mutation.predicates, ps...)
	return dtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dtd *DeviceTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dtd.hooks) == 0 {
		affected, err = dtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dtd.mutation = mutation
			affected, err = dtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dtd.hooks) - 1; i >= 0; i-- {
			mut = dtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtd *DeviceTypeDelete) ExecX(ctx context.Context) int {
	n, err := dtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dtd *DeviceTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: devicetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicetype.FieldID,
			},
		},
	}
	if ps := dtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dtd.driver, _spec)
}

// DeviceTypeDeleteOne is the builder for deleting a single DeviceType entity.
type DeviceTypeDeleteOne struct {
	dtd *DeviceTypeDelete
}

// Exec executes the deletion query.
func (dtdo *DeviceTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := dtdo.dtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{devicetype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dtdo *DeviceTypeDeleteOne) ExecX(ctx context.Context) {
	dtdo.dtd.ExecX(ctx)
}