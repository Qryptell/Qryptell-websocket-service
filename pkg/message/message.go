package message

type MessageType string

// Message types
const (
	SYSTEM_MSG MessageType = "SYSTEM_MSG"
	USER_MSG   MessageType = "USER_MSG"
	ACK_MSG    MessageType = "ACK_MSG"
)

type ContentType string

// Content Types
const (
	// for USER_MSG
	TEXT_MESSAGE ContentType = "TEXT_MESSAGE"
	FILE_MESSAGE ContentType = "FILE_MESSAGE"

	// for ACK_MSG
	MESSAGE_RECEIVED ContentType = "MESSAGE_RECEIVED"
	MESSAGE_DELIVERD ContentType = "MESSAGE_DELIVERD"
	MESSAGE_READ     ContentType = "MESSAGE_READ"

	// for SYSTEM_MSG
	CONNECTION_ID ContentType = "CONNECTION_ID"
)

// Websocket Message
type Msg struct {
	Id      string      `json:"id,omitempty"`
	From    string      `json:"from"`
	To      string      `json:"to,omitempty"`
	Type    MessageType `json:"type"`
	Time    string      `json:"time,omitempty"`
	Content ContentType `json:"content"`
	Message interface{} `json:"message"`
}

// Message for client
type ServerMsg struct {
	ConnectionId string `json:"connectionId,omitempty"`
	Msg          Msg    `json:"message"`
}
