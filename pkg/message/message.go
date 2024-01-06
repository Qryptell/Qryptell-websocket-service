package message

type MessageType string

// Message types
const (
	SYSTEM_MSG MessageType = "SYSTEM_MSG"
	USER_MSG   MessageType = "USER_MSG"
	ACK_MSG    MessageType = "ACK_MSG"
)

// Websocket Message
type Msg struct {
	Type    MessageType    `json:"type"`
	Message map[string]any `json:"message"`
}

// Message for client
type ClientMsg struct {
	ConnectionId string `json:"connectionId"`
	Msg          Msg    `json:"message"`
}
