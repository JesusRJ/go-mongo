package main

import (
	"context"

	"gitlab.com/brewnation/earth/pkg/core/arch"
	"gitlab.com/brewnation/earth/pkg/plugin/arch/repository/mongo"
	"gitlab.com/brewnation/earth/pkg/plugin/exp/slices"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type Repository[T any, E any] struct {
	provider arch.Repository[T]
	toMongo  arch.Mapper[E, T]
	toCore   arch.Mapper[T, E]
}

func NewRepository[T any, E any](coll *mongodb.Collection, mapFrom arch.Mapper[E, T], mapTo arch.Mapper[T, E]) *Repository[T, E] {
	return &Repository[T, E]{
		provider: mongo.NewRepository[T](coll),
		toMongo:  mapFrom,
		toCore:   mapTo,
	}
}

func (r *Repository[T, E]) Find(ctx context.Context, entity *E) (*E, error) {
	req, err := r.toMongo(entity)
	if err != nil {
		return nil, err
	}

	res, err := r.provider.Find(ctx, req)
	if err != nil {
		return nil, err
	}

	return r.toCore(res)
}

func (r *Repository[T, E]) Update(ctx context.Context, entity *E) (*E, error) {
	req, err := r.toMongo(entity)
	if err != nil {
		return nil, err
	}

	res, err := r.provider.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return r.toCore(res)
}

func (r *Repository[T, E]) Insert(ctx context.Context, entity *E) (*E, error) {
	req, err := r.toMongo(entity)
	if err != nil {
		return nil, err
	}

	res, err := r.provider.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return r.toCore(res)
}

func (r *Repository[T, E]) Delete(ctx context.Context, entity *E) error {
	req, err := r.toMongo(entity)
	if err != nil {
		return err
	}

	return r.provider.Delete(ctx, req)
}

func (r *Repository[T, E]) FindMany(ctx context.Context, entity *E) ([]E, error) {
	req, err := r.toMongo(entity)
	if err != nil {
		return nil, err
	}

	data, err := r.provider.FindMany(ctx, req)
	if err != nil {
		return nil, err
	}

	return slices.CloneByFunc(data, r.toCore), nil
}

func (r *Repository[T, E]) FindAll(ctx context.Context) ([]E, error) {
	data, err := r.provider.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return slices.CloneByFunc(data, r.toCore), nil
}
