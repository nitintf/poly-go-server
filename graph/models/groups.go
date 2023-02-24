package models

import "time"

type Groups struct {
	Id          string     `json:"id"`
	GroupName   string     `json:"group_name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
