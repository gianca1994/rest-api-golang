package models

import (
	"errors"
	"time"
)

const maxLenComment = 400

// Review models a review
type Review struct {
	Id      int64
	Stars   int
	Comment string
	Date    time.Time
}

// CreateReviewCMD create a new review from CMD
type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

// Validate validates the review
func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("invalid stars")
	}

	if len(cmd.Comment) > maxLenComment || len(cmd.Comment) < 1 {
		return errors.New("comment must be between 1 and 400 characters")
	}

	return nil
}
