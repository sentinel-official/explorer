package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	NodeStatisticCollectionName = "node_statistics"
)

func NodeStatisticFindOne(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.FindOneOptions) (bson.M, error) {
	var v bson.M
	if err := FindOne(ctx, db.Collection(NodeStatisticCollectionName), filter, &v, opts...); err != nil {
		return nil, findOneError(err)
	}

	return v, nil
}

func NodeStatisticInsertOne(ctx context.Context, db *mongo.Database, v bson.M, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return InsertOne(ctx, db.Collection(NodeStatisticCollectionName), v, opts...)
}

func NodeStatisticFindOneAndUpdate(ctx context.Context, db *mongo.Database, filter, update bson.M, opts ...*options.FindOneAndUpdateOptions) (bson.M, error) {
	var v bson.M
	if err := FindOneAndUpdate(ctx, db.Collection(NodeStatisticCollectionName), filter, update, &v, opts...); err != nil {
		return nil, findOneAndUpdateError(err)
	}

	return v, nil
}

func NodeStatisticFind(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.FindOptions) ([]bson.M, error) {
	var v []bson.M
	if err := Find(ctx, db.Collection(NodeStatisticCollectionName), filter, &v, opts...); err != nil {
		return nil, findError(err)
	}

	return v, nil
}

func NodeStatisticIndexesCreateMany(ctx context.Context, db *mongo.Database, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return IndexesCreateMany(ctx, db.Collection(NodeStatisticCollectionName), models, opts...)
}

func NodeStatisticInsertMany(ctx context.Context, db *mongo.Database, v bson.A, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return InsertMany(ctx, db.Collection(NodeStatisticCollectionName), v, opts...)
}

func NodeStatisticDeleteMany(ctx context.Context, db *mongo.Database, filter bson.M, opts ...*options.DeleteOptions) error {
	_, err := DeleteMany(ctx, db.Collection(NodeStatisticCollectionName), filter, opts...)
	if err != nil {
		return err
	}

	return nil
}

func NodeStatisticDrop(ctx context.Context, db *mongo.Database) error {
	return Drop(ctx, db.Collection(NodeStatisticCollectionName))
}

func NodeStatisticBulkWrite(ctx context.Context, db *mongo.Database, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return BulkWrite(ctx, db.Collection(NodeStatisticCollectionName), models, opts...)
}
