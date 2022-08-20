package gateway

import (
	"rest-api-golang/internal/database"
	"rest-api-golang/reviews/models"
)

type ReviewGateway interface {
	AddReview(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewGtw struct {
	ReviewStorage
}

func NewReviewGateway(mongoClient *database.Mongo) ReviewGateway {
	return &ReviewGtw{&ReviewStg{mongoClient}}
}