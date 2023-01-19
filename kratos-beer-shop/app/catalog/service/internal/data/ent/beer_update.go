// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-kratos/beer-shop/app/catalog/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/catalog/service/internal/data/ent/beer"
	"github.com/go-kratos/beer-shop/app/catalog/service/internal/data/ent/predicate"
)

// BeerUpdate is the builder for updating Beer entities.
type BeerUpdate struct {
	config
	hooks    []Hook
	mutation *BeerMutation
}

// Where adds a new predicate for the BeerUpdate builder.
func (bu *BeerUpdate) Where(ps ...predicate.Beer) *BeerUpdate {
	bu.mutation.predicates = append(bu.mutation.predicates, ps...)
	return bu
}

// SetName sets the "name" field.
func (bu *BeerUpdate) SetName(s string) *BeerUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetDescription sets the "description" field.
func (bu *BeerUpdate) SetDescription(s string) *BeerUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetCount sets the "count" field.
func (bu *BeerUpdate) SetCount(i int64) *BeerUpdate {
	bu.mutation.ResetCount()
	bu.mutation.SetCount(i)
	return bu
}

// AddCount adds i to the "count" field.
func (bu *BeerUpdate) AddCount(i int64) *BeerUpdate {
	bu.mutation.AddCount(i)
	return bu
}

// SetPrice sets the "price" field.
func (bu *BeerUpdate) SetPrice(i int64) *BeerUpdate {
	bu.mutation.ResetPrice()
	bu.mutation.SetPrice(i)
	return bu
}

// AddPrice adds i to the "price" field.
func (bu *BeerUpdate) AddPrice(i int64) *BeerUpdate {
	bu.mutation.AddPrice(i)
	return bu
}

// SetImages sets the "images" field.
func (bu *BeerUpdate) SetImages(b []biz.Image) *BeerUpdate {
	bu.mutation.SetImages(b)
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BeerUpdate) SetCreatedAt(t time.Time) *BeerUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BeerUpdate) SetNillableCreatedAt(t *time.Time) *BeerUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BeerUpdate) SetUpdatedAt(t time.Time) *BeerUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bu *BeerUpdate) SetNillableUpdatedAt(t *time.Time) *BeerUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// Mutation returns the BeerMutation object of the builder.
func (bu *BeerUpdate) Mutation() *BeerMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BeerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BeerUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BeerUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BeerUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BeerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   beer.Table,
			Columns: beer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: beer.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldName,
		})
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldDescription,
		})
	}
	if value, ok := bu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldCount,
		})
	}
	if value, ok := bu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldCount,
		})
	}
	if value, ok := bu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldPrice,
		})
	}
	if value, ok := bu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldPrice,
		})
	}
	if value, ok := bu.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: beer.FieldImages,
		})
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldCreatedAt,
		})
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{beer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BeerUpdateOne is the builder for updating a single Beer entity.
type BeerUpdateOne struct {
	config
	hooks    []Hook
	mutation *BeerMutation
}

// SetName sets the "name" field.
func (buo *BeerUpdateOne) SetName(s string) *BeerUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetDescription sets the "description" field.
func (buo *BeerUpdateOne) SetDescription(s string) *BeerUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetCount sets the "count" field.
func (buo *BeerUpdateOne) SetCount(i int64) *BeerUpdateOne {
	buo.mutation.ResetCount()
	buo.mutation.SetCount(i)
	return buo
}

// AddCount adds i to the "count" field.
func (buo *BeerUpdateOne) AddCount(i int64) *BeerUpdateOne {
	buo.mutation.AddCount(i)
	return buo
}

// SetPrice sets the "price" field.
func (buo *BeerUpdateOne) SetPrice(i int64) *BeerUpdateOne {
	buo.mutation.ResetPrice()
	buo.mutation.SetPrice(i)
	return buo
}

// AddPrice adds i to the "price" field.
func (buo *BeerUpdateOne) AddPrice(i int64) *BeerUpdateOne {
	buo.mutation.AddPrice(i)
	return buo
}

// SetImages sets the "images" field.
func (buo *BeerUpdateOne) SetImages(b []biz.Image) *BeerUpdateOne {
	buo.mutation.SetImages(b)
	return buo
}

// SetCreatedAt sets the "created_at" field.
func (buo *BeerUpdateOne) SetCreatedAt(t time.Time) *BeerUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BeerUpdateOne) SetNillableCreatedAt(t *time.Time) *BeerUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BeerUpdateOne) SetUpdatedAt(t time.Time) *BeerUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buo *BeerUpdateOne) SetNillableUpdatedAt(t *time.Time) *BeerUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// Mutation returns the BeerMutation object of the builder.
func (buo *BeerUpdateOne) Mutation() *BeerMutation {
	return buo.mutation
}

// Save executes the query and returns the updated Beer entity.
func (buo *BeerUpdateOne) Save(ctx context.Context) (*Beer, error) {
	var (
		err  error
		node *Beer
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BeerUpdateOne) SaveX(ctx context.Context) *Beer {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BeerUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BeerUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BeerUpdateOne) sqlSave(ctx context.Context) (_node *Beer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   beer.Table,
			Columns: beer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: beer.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Beer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldName,
		})
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldDescription,
		})
	}
	if value, ok := buo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldCount,
		})
	}
	if value, ok := buo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldCount,
		})
	}
	if value, ok := buo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldPrice,
		})
	}
	if value, ok := buo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldPrice,
		})
	}
	if value, ok := buo.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: beer.FieldImages,
		})
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldCreatedAt,
		})
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldUpdatedAt,
		})
	}
	_node = &Beer{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{beer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
