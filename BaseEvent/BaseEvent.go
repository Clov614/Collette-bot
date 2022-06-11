package BaseEvent

type MetaData struct {
	Time      int    `json:"time"`
	SelfID    int    `json:"self_id"`
	Post_type string `json:"post_type"`
}

type GeneralMsg struct {
	SubType     string `json:"sub_type"`
	MessageID   int    `json:"message_id"`
	UserId      int    `json:"user_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	MessageType string `json:"message_type"`
}

type MsgGroupEvent struct {
	MetaData
	GeneralMsg
	Anonymous  interface{} `json:"anonymous"`
	GroupID    int         `json:"group_id"`
	MessageSeq int         `json:"message_seq"`
	Sender     struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
}

type MsgPrivateEvent struct {
	MetaData
	GeneralMsg
	Sender struct {
		Age      int    `json:"age"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
	TargetID int `json:"target_id"`
}

type LoginEvent struct {
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfID        int    `json:"self_id"`
	SubType       string `json:"sub_type"`
	Time          int    `json:"time"`
}

type HeartBeat struct {
	Interval      int    `json:"interval"`
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfID        int    `json:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"PacketReceived"`
			PacketSent      int `json:"PacketSent"`
			PacketLost      int `json:"PacketLost"`
			MessageReceived int `json:"MessageReceived"`
			MessageSent     int `json:"MessageSent"`
			LastMessageTime int `json:"LastMessageTime"`
			DisconnectTimes int `json:"DisconnectTimes"`
			LostTimes       int `json:"LostTimes"`
		} `json:"stat"`
	} `json:"status"`
	Time int `json:"time"`
}

type PluginsData struct {
	Status  bool
	SendMsg string
}
