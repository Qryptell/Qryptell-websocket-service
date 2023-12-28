package message

import (
	"time"
)

type MessageType string
type AckStatus string

// Message types
const (
	SYSTEM_MSG MessageType = "SYSTEM_MSG"
	USER_MSG   MessageType = "USER_MSG"
	ACK_MSG    MessageType = "ACK_MSG"
)

// Ack status
const (
	RECEIVED AckStatus = "RECEIVED"
	READ     AckStatus = "READ"
	DELIVERD AckStatus = "DELIVERD"
)

// Websocket Message
type Msg struct {
	Type    MessageType    `json:"type"`
	Message map[string]any `json:"message"`
}

// User Message
type UserMsg struct {
	Id   string    `json:"messageId,omitempty"`
	From string    `json:"from"`
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

// Ack message send from client
type AckMessage struct {
	Id     string    `json:"messageId,omitempty"`
	From   string    `json:"from"`
	Status AckStatus `json:"status"`
	Time   time.Time `json:"Time"`
}
