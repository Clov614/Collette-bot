package sendMsgApi

type SENDGROUPMSG struct {
	Action string `json:"action"`
	Params struct {
		GroupID    int    `json:"group_id"`
		Message    string `json:"message"`
		AutoEscape bool   `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}

type SENDPRIVATEMSG struct {
	Action string `json:"action"`
	Params struct {
		UserID     int    `json:"user_id"`
		GroupID    string `json:"group_id"`
		Message    string `json:"message"`
		AutoEscape bool   `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}
