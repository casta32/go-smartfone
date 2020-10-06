package models

import (
	"errors"
	"time"
)

const MaxLengthComments = 400

//Review represent an anon review from website
type Review struct {
	Id      int64
	Starts  int       // 1 - 5
	Comment string    // max 400 chars
	Date    time.Time // created at
}

//CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Starts  int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Starts < 1 || cmd.Starts > 5 {
		return errors.New("stars must be between 1 - 5")
	}

	if len(cmd.Comment) > MaxLengthComments {
		return errors.New("comment must be less than 400 chars")
	}
	return nil
}
