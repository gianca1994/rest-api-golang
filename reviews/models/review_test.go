package models

import (
	"strings"
	"testing"
)

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{Stars: stars, Comment: comment}
}

func TestCreateReviewValidateCorrect(t *testing.T) {
	r := NewReview(4, "Iphone is the best")
	err := r.validate()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateReviewValidateFewerStars(t *testing.T) {
	r := NewReview(0, "Iphone is the best")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateReviewValidateMoreStars(t *testing.T) {
	r := NewReview(6, "Iphone is the best")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateReviewValidateWithoutComment(t *testing.T) {
	r := NewReview(5, "")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateReviewValidateWithExcessComment(t *testing.T) {
	r := NewReview(4, strings.Repeat("a", maxLenComment + 1))
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}



