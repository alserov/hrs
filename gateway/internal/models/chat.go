package models

import "time"

type Message struct {
	ID          string    `json:"id" db:"id"`
	Value       string    `json:"value" db:"value"`
	Files       []string  `json:"files" db:"-"`
	SenderID    string    `json:"senderID" db:"sender_id"`
	RecipientID string    `json:"recipientID" db:"recipient_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
