package models

type Todo struct {
	ID          string `json:"id,omitempty"  dynamo:"id,hash"`
	UserID      string `json:"user_id" dynamo:"user_id" index:"user_id,hash"`
	Title       string `json:"title" dynamo:"title"`
	IsCompleted bool   `json:"is_completed" dynamo:"is_completed"`
	CreatedAt   int64  `json:"created_at" dynamo:"created_at"`
	UpdatedAt   int64  `json:"updated_at" dynamo:"updated_at"`
}
