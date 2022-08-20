package gateway

import "rest-api-golang/reviews/models"

func (g *ReviewGtw) AddReview(cmd *models.CreateReviewCMD) (string, error) {
	return g.Add(cmd)
}
