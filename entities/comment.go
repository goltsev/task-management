package entities

import (
	"fmt"
	"time"
)

const (
	MinCommentText = 1
	MaxCommentText = 5000
)

type Comment struct {
	ID     int
	TaskID int
	Text   string
	Date   time.Time
}

func NewComment(text string) (*Comment, error) {
	c := &Comment{
		Text: text,
		Date: time.Now(),
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

// Validate checks Taskect Name and Description length
func (c *Comment) Validate() error {
	if len(c.Text) > MaxCommentText || len(c.Text) < MinCommentText {
		return fmt.Errorf("text has to be (%d-%d characters)",
			MinCommentText, MaxCommentText)
	}
	return nil
}
