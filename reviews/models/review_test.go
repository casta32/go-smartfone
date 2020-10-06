package models

import "testing"

func NewReview(starts int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Starts:  starts,
		Comment: comment,
	}
}

func Test_withCorrectParams(t *testing.T) {
	r := NewReview(4, "the iphone X looks good")
	err := r.validate()

	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()
	}
}

func Test_shouldFailWhitWrongNumerOfStars(t *testing.T) {
	r := NewReview(8, "the iphone X looks good")
	err := r.validate()

	if err == nil {
		t.Error("should fail whit 8 stars")
		t.Fail()
	}
}
