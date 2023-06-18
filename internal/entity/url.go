package entity

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	missingOriginalURLError = "original url must be provided"
	invalidOriginalURLError = "original url is invalid"
	missingShortURLError    = "short url must be provided"
	invalidShortURLError    = "short url is invalid"
)

type URLEntity struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	UserID        primitive.ObjectID `json:"-" bson:"user_id"`
	OriginalUrl   string             `json:"original_url" bson:"original_url"`
	ShortUrl      string             `json:"short_url" bson:"short_url"`
	TotalAccesses uint64             `json:"-" bson:"total_accesses"`
	TotalReports  uint64             `json:"-" bson:"total_reports"`
	Reputation    float64            `json:"reputation" bson:"-"`
	IsActived     bool               `json:"is_actived" bson:"is_actived"`
	IsDeleted     bool               `json:"-" bson:"is_deleted"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at"`
}

func NewURLEntity(originalURL, shortURL string) *URLEntity {
	return &URLEntity{
		OriginalUrl: originalURL,
		ShortUrl:    shortURL,
		IsActived:   true,
	}
}

func (u *URLEntity) OriginalURLValidator() error {
	if u.OriginalUrl == "" {
		return errors.New(missingOriginalURLError)
	}

	if !isUrl(u.OriginalUrl) {
		return errors.New(invalidOriginalURLError)
	}

	u.OriginalUrl = strings.ToLower(u.OriginalUrl)

	return nil
}

func (u *URLEntity) ShortURLValidator() error {
	if u.ShortUrl == "" {
		return errors.New(missingShortURLError)
	}

	if !isUrl(u.ShortUrl) {
		return errors.New(invalidShortURLError)
	}

	return nil
}
