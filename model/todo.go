package model

import (
	"time"
)

type Todo struct {
	Id int `json:"id"`
	Content string `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	ModifiedAt time.Time `json:"modifiedAt,omitempty"`
}