package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findOneError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}

	return err
}

func findOneAndUpdateError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}

	return err
}

func findError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}

	return err
}

func FindOne(ctx context.Context, c *mongo.Collection, filter bson.M, v interface{}, opts ...*options.FindOneOptions) error {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "FindOne", time.Since(now))
	}()

	result := c.FindOne(ctx, filter, opts...)
	if result.Err() != nil {
		return result.Err()
	}
	if err := result.Decode(v); err != nil {
		return err
	}

	return nil
}

func InsertOne(ctx context.Context, c *mongo.Collection, v interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "InsertOne", time.Since(now))
	}()

	return c.InsertOne(ctx, v, opts...)
}

func InsertMany(ctx context.Context, c *mongo.Collection, v []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "InsertMany", time.Since(now))
	}()

	return c.InsertMany(ctx, v, opts...)
}

func FindOneAndUpdate(ctx context.Context, c *mongo.Collection, filter, update bson.M, v interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "FindOneAndUpdate", time.Since(now))
	}()

	result := c.FindOneAndUpdate(ctx, filter, update, opts...)
	if result.Err() != nil {
		return result.Err()
	}
	if err := result.Decode(v); err != nil {
		return err
	}

	return nil
}

func Find(ctx context.Context, c *mongo.Collection, filter bson.M, v interface{}, opts ...*options.FindOptions) error {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "Find", time.Since(now))
	}()

	cursor, err := c.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, v); err != nil {
		return err
	}

	return nil
}

func Aggregate(ctx context.Context, c *mongo.Collection, pipeline []bson.M, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "Aggregate", time.Since(now))
	}()

	return c.Aggregate(ctx, pipeline, opts...)
}

func AggregateAll(ctx context.Context, c *mongo.Collection, pipeline []bson.M, v interface{}, opts ...*options.AggregateOptions) error {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "AggregateAll", time.Since(now))
	}()

	cursor, err := c.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, v); err != nil {
		return err
	}

	return nil
}

func CountDocuments(ctx context.Context, c *mongo.Collection, filter bson.M, opts ...*options.CountOptions) (int64, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "CountDocuments", time.Since(now))
	}()

	return c.CountDocuments(ctx, filter, opts...)
}

func Distinct(ctx context.Context, c *mongo.Collection, fieldName string, filter bson.M, opts ...*options.DistinctOptions) (bson.A, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "Distinct", time.Since(now))
	}()

	return c.Distinct(ctx, fieldName, filter, opts...)
}

func Drop(ctx context.Context, c *mongo.Collection) error {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "Drop", time.Since(now))
	}()

	return c.Drop(ctx)
}

func DeleteMany(ctx context.Context, c *mongo.Collection, filter bson.M, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "DeleteMany", time.Since(now))
	}()

	return c.DeleteMany(ctx, filter, opts...)
}

func UpdateMany(ctx context.Context, c *mongo.Collection, filter, update bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "UpdateMany", time.Since(now))
	}()

	return c.UpdateMany(ctx, filter, update, opts...)
}

func IndexesCreateMany(ctx context.Context, c *mongo.Collection, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "IndexesCreateMany", time.Since(now))
	}()

	return c.Indexes().CreateMany(ctx, models, opts...)
}

func BulkWrite(ctx context.Context, c *mongo.Collection, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	now := time.Now()
	defer func() {
		log.Println(c.Name(), "BulkWrite", time.Since(now))
	}()

	return c.BulkWrite(ctx, models, opts...)
}
