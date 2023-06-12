package repository

import (
	"context"
	"errors"
	"time"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) IsEmailRegistered(ctx context.Context, email string) bool {
	client := NewMongoConnection(ctx)
	defer client.Disconnect(ctx)

	coll := client.Database("url_shortener").Collection("users")
	err := coll.FindOne(
		ctx,
		bson.D{primitive.E{Key: "email", Value: email}},
	).Err()

	return err == nil
}

func (u UserRepository) CreateUser(ctx context.Context, user *entity.UserEntity) {
	client := NewMongoConnection(ctx)
	defer client.Disconnect(ctx)

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		users_coll := client.Database("url_shortener").Collection("users")
		users_details_coll := client.Database("url_shortener").Collection("user_details")

		// Insert User
		result, err := users_coll.InsertOne(ctx, bson.D{
			{Key: "name", Value: user.Name},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
			{Key: "created_at", Value: time.Now()},
			{Key: "updated_at", Value: time.Now()},
		})
		if err != nil {
			panic(err)
		}

		// Map user id
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			user.ID = oid
			user.UserDetail.UserID = oid
		} else {
			panic(errors.New("error to get inserted user id"))
		}

		// Insert User Detail
		if _, err = users_details_coll.InsertOne(ctx, bson.D{
			{Key: "user_id", Value: primitive.ObjectID(user.UserDetail.UserID)},
			{Key: "all_accesses", Value: user.UserDetail.AllAccesses},
			{Key: "all_reports", Value: user.UserDetail.AllReports},
			{Key: "updated_at", Value: time.Now()},
		}); err != nil {
			panic(err)
		}

		if err := session.CommitTransaction(ctx); err != nil {
			panic(err)
		}

		return nil, nil
	})

	if err != nil {
		panic(err)
	}
}

func (u UserRepository) LoadUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	client := NewMongoConnection(ctx)
	defer client.Disconnect(ctx)

	var user entity.UserEntity
	users_coll := client.Database("url_shortener").Collection("users")
	if err := users_coll.FindOne(
		ctx,
		bson.D{primitive.E{Key: "email", Value: email}},
	).Decode(&user); err != nil {
		return nil, err
	}
	users_detail_coll := client.Database("url_shortener").Collection("user_details")
	if err := users_detail_coll.FindOne(
		ctx,
		bson.D{primitive.E{Key: "user_id", Value: user.ID}},
	).Decode(&user.UserDetail); err != nil {
		return nil, err
	}

	return &user, nil
}
