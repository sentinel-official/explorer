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

func NodeInsertOne(ctx context.Context, db *mongo.Database, v *types.Node, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return InsertOne(ctx, db.Collection(NodeCollectionName), v, opts...)
}

func NodeFindOneAndUpdate(ctx context.Context, db *mongo.Database, filter, update bson.M, opts ...*options.FindOneAndUpdateOptions) (*types.Node, error) {
	var v types.Node
	if err := FindOneAndUpdate(ctx, db.Collection(NodeCollectionName), filter, update, &v, opts...); err != nil {
		return nil, findOneAndUpdateError(err)
	}

	return &v, nil
}

func NodeFind(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.FindOptions) ([]*types.Node, error) {
	var v []*types.Node
	if err := Find(ctx, db.Collection(NodeCollectionName), filter, &v, opts...); err != nil {
		return nil, findError(err)
	}

	return v, nil
}