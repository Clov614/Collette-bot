package eventListener

import (
	"Collette_bot/eventListener/middleHandler"
	"Collette_bot/network/ws"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	loginevent LoginEvent
)

// 监听event
func Listen(Event []byte, hub *ws.Hub) {
	loginEvent(Event, hub)
	msgGroupEvent(Event, hub)
	msgPrivateEvent(Event, hub)
}

// 监听登录事件
func loginEvent(Event []byte, hub *ws.Hub) {
	err := json.Unmarshal(Event, &loginevent)
	if err != nil {
		return
	}
	// 忽略心跳事件
	if loginevent.MetaEventType != "lifecycle" { // 忽略心跳事件
		return
	}
	if loginevent.MetaEventType == "heartbeat" { // 忽略心跳事件
		return
	}
	//hub.LoginSuccess <- true
	log.Info(fmt.Sprintf("QQ: %v 连接gocqhttp成功！", loginevent.SelfID))
	return
}

// 监听群组消息事件
func msgGroupEvent(Event []byte, hub *ws.Hub) {
	//log.Info(string(Event))
	msgevent := MsgGroupEvent{}
	err := json.Unmarshal(Event, &msgevent)
	if err != nil {
		log.Error(err)
	}
	//log.Info(msgevent.MessageType)
	if msgevent.MessageType == "group" {
		// 替换特殊字符
		newMsg := ChangeSpecialsymbols(msgevent.Message)
		msgevent.Message = newMsg
		log.Println("收到群组消息：", msgevent.UserId, msgevent.Sender.Nickname, msgevent.Message)
		// status 为是否发送消息的flag
		status, sendMsg := middleHandler.PostGROUPmsg(msgevent.GroupID, msgevent.Message)
		if status == true {
			hub.Sendmsg <- sendMsg
		}
	}
	return
}

// 监听私聊消息
func msgPrivateEvent(Event []byte, hub *ws.Hub) {
	msgevent := MsgPrivateEvent{}
	err := json.Unmarshal(Event, &msgevent)
	if err != nil {
		log.Error(err)
	}

	if msgevent.MessageType == "private" {
		// 替换特殊字符
		newMsg := ChangeSpecialsymbols(msgevent.Message)
		msgevent.Message = newMsg
		log.Println("收到私聊消息：", msgevent.UserId, msgevent.Sender.Nickname, msgevent.Message)
		status, sendMsg := middleHandler.PostPRIVATEmsg(msgevent.UserId, msgevent.Message)
		if status == true {
			hub.Sendmsg <- sendMsg
		}
	}

	return

}

// 转换特殊符号
func ChangeSpecialsymbols(oldMessage string) (newMessage string) {
	var (
		commaSymb     = "&#44;"
		comma         = ","
		etSymb        = "&amp;"
		et            = "&"
		fbracketSymb  = "&#91;"
		fbracket      = "["
		bebracketSymb = "&#93;"
		bebracket     = "]"
	)
	newMessage = strings.Replace(oldMessage, commaSymb, comma, -1)
	newMessage = strings.Replace(newMessage, etSymb, et, -1)
	newMessage = strings.Replace(newMessage, fbracketSymb, fbracket, -1)
	newMessage = strings.Replace(newMessage, bebracketSymb, bebracket, -1)
	return
}
