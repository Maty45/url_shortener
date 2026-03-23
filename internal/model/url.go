package model

import "time"

type URL struct {
	Code      string
	Original  string
	Clicks    int
	CreatedAt time.Time
}
