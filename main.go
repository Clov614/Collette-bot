package main

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/Subs"
	"Collette_bot/eventListener"
	_ "Collette_bot/log"
	"Collette_bot/network/ws"
	_ "Collette_bot/setting"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go func() {
		wg.Add(1)
		defer wg.Done()
		http.HandleFunc("/universal/", ws.WsHandle)
		err := http.ListenAndServe(":6700", nil)

		if err != nil {
			log.Error("断开ws连接\n")
			log.Fatalln(err)
		}
		wg.Wait()
	}()

	log.Info(fmt.Sprintf("正向ws启动，监听于 127.0.0.1:6700 ...."))
	hub := ws.Hhub

	// 消息监听器
	go func() {
		wg.Add(1)
		for v := range hub.Chmsg {
			HeartbeatFilter(string(v))
			eventListener.Listen(v, hub)
		}
	}()

	// 消息订阅器
	go func() {
		for {
			Subs.SubscribeHandle(hub)
		}
	}()
	// 消息发送器
	go func() {
		for v := range hub.Sendmsg {
			log.Info(v)
			conn := ws.GetClient()
			_ = conn.WriteJSON(v)

		}
	}()
	wg.Wait()
}

var metaData BaseEvent.MetaData

// 心跳事件过滤
func HeartbeatFilter(str string) {

	re, _ := regexp.Compile("{\"interval\":5000,\"meta_event_type\":\"heartbeat\"")
	if !re.MatchString(str) {
		// 替换特殊字符
		str = eventListener.ChangeSpecialsymbols(str)
		//log.Info(str)
		json.Unmarshal([]byte(str), &metaData)
		// 测试使用，查看Post_type
		//log.Info(metaData.Post_type)
	}
}
