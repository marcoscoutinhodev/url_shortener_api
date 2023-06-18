package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepository struct{}

func NewURLRepository() *URLRepository {
	return &URLRepository{}
}

func (u URLRepository) CreateShortURL(ctx context.Context, url *entity.URLEntity, userId string) {
	client := NewMongoConnection(ctx)
	defer client.Disconnect(ctx)

	userIdAsObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		panic(err)
	}

	url.UserID = userIdAsObjectID

	urls_coll := client.Database("url_shortener").Collection("urls")
	if _, err = urls_coll.InsertOne(ctx, bson.D{
		{Key: "original_url", Value: url.OriginalUrl},
		{Key: "short_url", Value: url.ShortUrl},
		{Key: "total_accesses", Value: url.TotalAccesses},
		{Key: "total_reports", Value: url.TotalReports},
		{Key: "is_actived", Value: url.IsActived},
		{Key: "is_deleted", Value: url.IsDeleted},
		{Key: "user_id", Value: url.UserID},
		{Key: "created_at", Value: time.Now()},
	}); err != nil {
		panic(err)
	}
}

func (u URLRepository) GetOriginalURL(ctx context.Context, shortURL string) (*entity.URLEntity, error) {
	client := NewMongoConnection(ctx)
	defer client.Disconnect(ctx)

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	var url entity.URLEntity
	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {

		urls_coll := client.Database("url_shortener").Collection("urls")
		if err := urls_coll.FindOne(
			ctx,
			bson.D{
				primitive.E{Key: "short_url", Value: shortURL},
				primitive.E{Key: "is_deleted", Value: false},
			},
		).Decode(&url); err != nil {
			return nil, err
		}

		url.TotalAccesses += 1

		if _, err := urls_coll.UpdateOne(
			ctx,
			bson.D{primitive.E{Key: "_id", Value: url.ID}},
			bson.M{"$set": bson.M{"total_accesses": url.TotalAccesses}},
		); err != nil {
			return nil, err
		}

		if err := session.CommitTransaction(ctx); err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &url, nil
}

func (u URLRepository) ReportURL(ctx context.Context, urlID string) error {
	client := NewMongoConnection(ctx)
	defer client.Disconnect(ctx)

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	urlIDAsObjectID, err := primitive.ObjectIDFromHex(urlID)
	if err != nil {
		return errors.New("invalid url id")
	}

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		var url entity.URLEntity

		urls_coll := client.Database("url_shortener").Collection("urls")
		if err := urls_coll.FindOne(
			ctx,
			bson.D{
				primitive.E{Key: "_id", Value: urlIDAsObjectID},
			},
		).Decode(&url); err != nil {
			return nil, err
		}

		url.TotalReports += 1

		if _, err := urls_coll.UpdateOne(
			ctx,
			bson.D{primitive.E{Key: "_id", Value: url.ID}},
			bson.M{"$set": bson.M{"total_reports": url.TotalReports}},
		); err != nil {
			return nil, err
		}

		if err := session.CommitTransaction(ctx); err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return errors.New("no matching url")
	}

	return nil
}
