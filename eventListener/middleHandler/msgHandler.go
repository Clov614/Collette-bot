package middleHandler

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/func_module"
	SendAPI "Collette_bot/sendMsgApi"
)

var (
	sendgroupMsg   SendAPI.SENDGROUPMSG
	sendprivateMsg SendAPI.SENDPRIVATEMSG
	pluginsMsg     BaseEvent.PluginsMsg
)

// 将消息处理为待发送JSON_struct  内层判断消息以及回复消息 返回sendMsgApi.SENDGROUPMSG
func PostGROUPmsg(msgEvent BaseEvent.MsgGroupEvent) (bool, SendAPI.SENDGROUPMSG) {
	sendgroupMsg.Action = "send_group_msg"
	sendgroupMsg.Params.GroupID = msgEvent.GroupID
	sendgroupMsg.Echo = ""
	sendgroupMsg.Params.AutoEscape = false
	// 拟定使用通道进行消息的判断，将功能再度抽象到外层
	// 使用Done来进行返回值的通信
	//send_message := make(chan string)
	// 组合Meta General message =
	pluginsMsg.GeneralMsg = msgEvent.GeneralMsg
	pluginsMsg.MetaData = msgEvent.MetaData
	// 交由功能判断器进行处理
	done, message := func_module.PluginsDetermine(pluginsMsg)
	if done {
		sendgroupMsg.Params.Message = message
		return true, sendgroupMsg
	} else {
		return false, SendAPI.SENDGROUPMSG{} // 不满足判断条件返回
	}
}

func PostPRIVATEmsg(msgEvent BaseEvent.MsgPrivateEvent) (bool, SendAPI.SENDPRIVATEMSG) {
	sendprivateMsg.Action = "send_private_msg"
	sendprivateMsg.Params.UserID = msgEvent.UserId
	sendprivateMsg.Echo = ""
	sendprivateMsg.Params.AutoEscape = false
	// 组合Meta General message =
	pluginsMsg.GeneralMsg = msgEvent.GeneralMsg
	pluginsMsg.MetaData = msgEvent.MetaData
	// 交由功能判断器进行处理
	done, message := func_module.PluginsDetermine(pluginsMsg)
	if done {
		sendprivateMsg.Params.Message = message
		return true, sendprivateMsg
	} else {
		return false, SendAPI.SENDPRIVATEMSG{}
	}

}
