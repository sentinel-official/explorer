package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sentinel-official/explorer/types"
)

const (
	NodeCollectionName = "nodes"
)

func NodeFindOne(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.FindOneOptions) (*types.Node, error) {
	var v types.Node
	if err := FindOne(ctx, db.Collection(NodeCollectionName), filter, &v, opts...); err != nil {
		return nil, findOneError(err)
	}

	return &v, nil
}

func NodeSave(ctx context.Context, db *mongo.Database, v *types.Node, opts ...*options.InsertOneOptions) error {
	return Save(ctx, db.Collection(NodeCollectionName), v, opts...)
}

func NodeFindOneAndUpdate(ctx context.Context, db *mongo.Database, filter, update bson.M, opts ...*options.FindOneAndUpdateOptions) (*types.Node, error) {
	var v types.Node
	if err := FindOneAndUpdate(ctx, db.Collection(NodeCollectionName), filter, update, &v, opts...); err != nil {
		return nil, findOneAndUpdateError(err)
	}

	return &v, nil
}

func NodeFindAll(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.FindOptions) ([]*types.Node, error) {
	var v []*types.Node
	if err := FindAll(ctx, db.Collection(NodeCollectionName), filter, &v, opts...); err != nil {
		return nil, findError(err)
	}

	return v, nil
}

func NodeAggregate(ctx context.Context, db *mongo.Database, pipeline []bson.M, opts ...*options.AggregateOptions) ([]bson.M, error) {
	var v []bson.M
	if err := Aggregate(ctx, db.Collection(NodeCollectionName), pipeline, &v, opts...); err != nil {
		return nil, findError(err)
	}

	return v, nil
}

func NodeCountDocuments(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.CountOptions) (int64, error) {
	return CountDocuments(ctx, db.Collection(NodeCollectionName), filter, opts...)
}
