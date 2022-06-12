package ping

import (
	"Collette_bot/BaseEvent"
)

func Ping(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData) {
	dataCheck.SendMsg = "pong!"
	switch msgEvent.Message {
	case "ping":
		dataCheck.Status = true
	case "Ping":
		dataCheck.Status = true
	}
	return
}
