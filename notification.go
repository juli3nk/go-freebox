package freebox

type WebSocketNotification struct {
	Action  string      `json:"action"`
	Success bool        `json:"success"`
	Source  string      `json:"source,omitempty"`
	Event   string      `json:"event,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
