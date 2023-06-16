package repository

import (
	"context"
	"time"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
