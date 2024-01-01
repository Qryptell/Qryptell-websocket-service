package message 

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
