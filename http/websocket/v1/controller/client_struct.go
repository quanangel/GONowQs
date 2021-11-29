package controller

type _base struct {
	Action string `json:"action"`
}

type _Login struct {
	_base
	UserType string `json:"user_type"`
	Token    string `json:"token"`
}

// Message is response message struct
type Message struct {
	Action string                 `json:"action"`
	Code   int                    `json:"code"`
	Msg    string                 `json:"msg"`
	Data   map[string]interface{} `json:"data"`
}
