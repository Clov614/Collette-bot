package ping

import (
	"Collette_bot/BaseEvent"
)

func Ping(msgEvent BaseEvent.GeneralMsg, dataCheck *BaseEvent.PluginsData) {
	dataCheck.SendMsg = "pong!"
	switch msgEvent.Message {
	case "ping":
		dataCheck.Status = true
	case "Ping":
		dataCheck.Status = true
	}
	return
}
