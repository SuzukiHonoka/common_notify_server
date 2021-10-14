package api

type Reply struct {
	Action string      `json:"action"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func NewReply(action string, status bool, data interface{}) *Reply {
	return &Reply{
		Action: action,
		Status: status,
		Data:   data,
	}
}
