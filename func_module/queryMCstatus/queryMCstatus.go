package queryMCstatus

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/setting"
	"fmt"
	"github.com/PassTheMayo/mcstatus/v3"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func QuerymcStatus(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData) {
	receiveMsg := msgEvent.Message
	// 获取时间
	tm := time.Unix(int64(msgEvent.Time), 0)
	Time := tm.Format("2006-01-02 15:04:05")
	// 判断是否有at机器人的变量
	var ATstatus bool
	re, _ := regexp.Compile("\\[CQ:at,qq=" + setting.Data.SelfQQ + "]")
	// 判断是否有at机器人
	for _, V := range setting.Data.Nickname {
		reNick, _ := regexp.Compile(V)
		if reNick.MatchString(receiveMsg) {
			ATstatus = true
		}
	}
	if re.MatchString(receiveMsg) {
		ATstatus = true
	}
	if ATstatus {
		targets := []string{"服务器状态", "服务器查询", "MC服务器", "mc服务器"}
		for _, v := range targets {
			reServer, _ := regexp.Compile(v)
			if reServer.MatchString(receiveMsg) {
				dataCheck.Status = true
				dataCheck.SendMsg = queryStatus(Time)
			}
		}

	}
}

func queryStatus(time string) (sendMsg string) {
	serverPosit := "mc.taiga.icu"
	response, err := mcstatus.Status(serverPosit, 25565)

	if err != nil {
		log.Error(err)
	}

	// 拼接返回信息
	sendMsg = fmt.Sprintf("当前服务器状态: \n 服务器: %s \n 游戏版本: %s \n协议: %d \n", serverPosit, response.Version.Name, response.Version.Protocol)
	sendMsg += fmt.Sprintf("在线人数: %d  最大人数: %d\n", response.Players.Online, response.Players.Max)
	sendMsg += fmt.Sprintf("在线玩家: \n")
	for _, Sample := range response.Players.Sample {
		sendMsg += fmt.Sprintf("%s \n", Sample.Name)
	}
	sendMsg += "当前时间: " + time
	return sendMsg
}
