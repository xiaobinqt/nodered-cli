package model

type FlowItem struct {
	Id       string        `json:"id"`
	Type     string        `json:"type"`
	Label    string        `json:"label,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
	Info     string        `json:"info,omitempty"`
	Env      []interface{} `json:"env,omitempty"`
	Z        string        `json:"z,omitempty"`
	Name     string        `json:"name,omitempty"`
	Props    []struct {
		P  string `json:"p"`
		Vt string `json:"vt,omitempty"`
	} `json:"props,omitempty"`
	Repeat      string          `json:"repeat,omitempty"`
	Crontab     string          `json:"crontab,omitempty"`
	Once        bool            `json:"once,omitempty"`
	OnceDelay   float64         `json:"onceDelay,omitempty"`
	Topic       string          `json:"topic,omitempty"`
	Payload     string          `json:"payload,omitempty"`
	PayloadType string          `json:"payloadType,omitempty"`
	X           int             `json:"x,omitempty"`
	Y           int             `json:"y,omitempty"`
	Wires       [][]interface{} `json:"wires,omitempty"`
	Links       []interface{}   `json:"links,omitempty"`
	Timeout     string          `json:"timeout,omitempty"`
}
