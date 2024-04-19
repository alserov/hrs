package models

import "time"

type Message struct {
	ID          string    `json:"id"`
	Value       string    `json:"value"`
	Files       []string  `json:"files"`
	SenderID    string    `json:"senderID"`
	RecipientID string    `json:"recipientID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type MessageEdit struct {
	MessageID string    `json:"messageID"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updatedAt"`
}
