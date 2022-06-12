package eventListener

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/eventListener/middleHandler"
	"Collette_bot/network/ws"
	"Collette_bot/setting"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	loginevent BaseEvent.LoginEvent
)

// 监听event
func Listen(Event []byte, hub *ws.Hub) {
	loginEvent(Event, hub)
	messageEvent(Event, hub)
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
	setting.Data.SelfQQ = strconv.Itoa(loginevent.SelfID)
	return
}

// 处理群组消息事件
func msgGroupEvent(Event []byte, hub *ws.Hub) {
	msgevent := BaseEvent.MsgGroupEvent{}
	err := json.Unmarshal(Event, &msgevent)
	if err != nil {
		log.Error(err)
	}
	// 替换特殊字符
	newMsg := ChangeSpecialsymbols(msgevent.Message)
	msgevent.Message = newMsg
	log.Printf("收到群组消息 ID:[%d] Name:[%s] Msg: %s", msgevent.UserId, msgevent.Sender.Nickname, msgevent.Message)
	// status 为是否发送消息的flag
	status, sendMsg := middleHandler.PostGROUPmsg(msgevent)
	if status == true {
		hub.Sendmsg <- sendMsg
	}
	return
}

// 处理私聊消息
func msgPrivateEvent(Event []byte, hub *ws.Hub) {
	msgevent := BaseEvent.MsgPrivateEvent{}
	err := json.Unmarshal(Event, &msgevent)
	if err != nil {
		log.Error(err)
	}
	// 替换特殊字符
	newMsg := ChangeSpecialsymbols(msgevent.Message)
	msgevent.Message = newMsg
	log.Printf("收到私聊消息 ID:[%d] Name:[%s] Msg: %s", msgevent.UserId, msgevent.Sender.Nickname, msgevent.Message)
	status, sendMsg := middleHandler.PostPRIVATEmsg(msgevent)
	if status == true {
		hub.Sendmsg <- sendMsg
	}
}

// 监听消息
func messageEvent(Event []byte, hub *ws.Hub) {
	msgEvent := BaseEvent.GeneralMsg{}
	err := json.Unmarshal(Event, &msgEvent)
	if err != nil {
		log.Error(err)
	}
	if msgEvent.MessageType == "group" {
		msgGroupEvent(Event, hub)
	}
	if msgEvent.MessageType == "private" {
		msgPrivateEvent(Event, hub)
	}
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
