package models

import "time"

type HistoryParam struct {
	UserID      string `json:"userID"`
	RecipientID string `json:"recipientID"`
	Offset      int32  `json:"offset"`
	Limit       int32  `json:"limit"`
}

type ChatsParam struct {
	UserID string `json:"userID"`
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
}

type Chat struct {
	LastMessage Message `json:"lastMessage"`
}

type Message struct {
	ID          string    `json:"id" db:"id"`
	Value       string    `json:"value" db:"value"`
	Files       []string  `json:"files" db:"-"`
	SenderID    string    `json:"senderID" db:"sender_id"`
	RecipientID string    `json:"recipientID" db:"recipient_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type MessageEdit struct {
	MessageID string    `json:"messageID"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updatedAt"`
}
