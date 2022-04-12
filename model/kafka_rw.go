package model

type KafkaRwOrder struct {
	Method string  `json:"method,omitempty"`
	Body   []Order `json:"body,omitempty"`
	Id     int     `json:"id,omitempty"`
}
