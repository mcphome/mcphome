package mcp

type JSONRPC struct {
	Version string `json:"jsonrpc"`
	ID      string `json:"id"`
}

type Request struct {
	Method string         `json:"method"`
	Params map[string]any `json:"params"`
}

type Notification struct {
	Method string         `json:"method"`
	Params map[string]any `json:"params"`
}

type Result struct {
	Result map[string]any `json:"result"`
	Error  Error          `json:"error"`
}

type Error struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

type TransportType int

const (
	TransportTypeStdio TransportType = iota
	TransportTypeHTTP
)
