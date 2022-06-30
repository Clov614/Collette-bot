package Subs

import (
	"Collette_bot/network/ws"
	SendAPI "Collette_bot/sendMsgApi"
	"Collette_bot/setting"
	"fmt"
	"github.com/PassTheMayo/mcstatus/v3"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// 订阅消息基本结构
//type SubBaseInfo struct {
//	StartTime       int64                  `json:"startTime"`
//	Interval        int64                  `json:"interval"`
//	SubName         []string               `json:"subName"`
//	SubscriptionObj map[string]interface{} `json:"subscriptionObj"`
//}

// 订阅MC服务器信息
type SubsMcSrvInfo struct {
	UserIds   []int               `json:"user_ids"`
	GroupIDs  []int               `json:"group_ids"`
	AddrSrv   []string            `json:"addr_srv"`
	Interval  int64               `json:"interval"`
	TempTime  int64               `json:"tempTime"`
	TempInfos map[string]TempInfo `json:"temp_infos"`
}

type TempInfo struct {
	Players int `json:"players"`
}

var (
	sendgroupMsg   SendAPI.SENDGROUPMSG
	sendprivateMsg SendAPI.SENDPRIVATEMSG
	//pluginsMsg     BaseEvent.PluginsMsg
	//SubBaseinfo SubBaseInfo
	SubMC SubsMcSrvInfo
)

func init() {
	//SubBaseinfo = SubBaseInfo{
	//	StartTime: time.Now().Unix(),
	//	Interval:  60, // 默认订阅心跳为一分钟(节省性能开销)
	//}
	if !setting.PathExists("./Source/SubMC.yml") {
		SubMC = SubsMcSrvInfo{
			Interval: 60,
			TempTime: time.Now().Unix(),
			AddrSrv:  []string{"mc.taiga.icu"},
		}
		setting.WriteYaml(SubMC, "./Source/SubMC.yml")
	}
	setting.ReadYaml(&SubMC, "./Source/SubMC.yml")
}

// 处理订阅服务
func SubscribeHandle(hub *ws.Hub) {
	SubMC.SubsMcServerInfo(hub)
}

func (SubMc *SubsMcSrvInfo) SubsMcServerInfo(hub *ws.Hub) {
	// 设置订阅更新间隔
	heartBeat := SubMc.Interval
	heartBeatD := time.Duration(heartBeat) * time.Second
	time.Sleep(heartBeatD)

	// 读取
	setting.ReadYaml(&SubMC, "./Source/SubMC.yml")
	// 设置条件满足更新订阅

	// 遍历每个服务器地址
	for _, v := range SubMc.AddrSrv {
		// 防止Players为空
		if strconv.Itoa(SubMC.TempInfos[v].Players) == "" {
			SubMc.TempInfos[v] = TempInfo{
				Players: 0,
			}
		}

		// 获取到Mc服务器信息，以及当前玩家数量

		mcInfoMsg, nowPlayer, err := queryMcStatus(v, SubMc.TempInfos[v].Players)
		if err != nil {
			return
		}
		if nowPlayer-SubMc.TempInfos[v].Players != 0 {
			SubMc.TempInfos[v] = TempInfo{
				Players: nowPlayer,
			}
			setting.WriteYaml(&SubMc, "./Source/SubMC.yml")
			for _, v := range SubMc.UserIds {
				sendMsg("private", hub, mcInfoMsg, v, 0)
			}
			for _, v := range SubMc.GroupIDs {
				sendMsg("group", hub, mcInfoMsg, 0, v)
			}
		}
	}

}

func queryMcStatus(srvAddr string, oldPlayer int) (sendMsg string, nowPlayer int, err error) {
	response, err := mcstatus.Status(srvAddr, 25565)
	if err != nil {
		log.Error(err)
		return "", 0, err
	}
	nowPlayer = response.Players.Online
	// 拼接返回信息
	sendMsg = fmt.Sprintf("当前服务器状态: \n 服务器: %s \n 游戏版本: %s \n协议: %d \n", srvAddr, response.Version.Name, response.Version.Protocol)
	sendMsg += fmt.Sprintf("在线人数: %d (%d)  最大人数: %d\n", response.Players.Online, response.Players.Online-oldPlayer, response.Players.Max)
	sendMsg += fmt.Sprintf("在线玩家: \n")
	for _, Sample := range response.Players.Sample {
		sendMsg += fmt.Sprintf("%s \n", Sample.Name)
	}
	sendMsg += "当前时间: " + strconv.FormatInt(time.Now().Unix(), 10)
	return sendMsg, nowPlayer, nil
}

func sendMsg(_msgType string, hub *ws.Hub, sendMsg string, userId int, groupId int) {
	if _msgType == "group" {
		sendgroupMsg.Action = "send_group_msg"
		sendgroupMsg.Params.GroupID = groupId
		sendgroupMsg.Echo = ""
		sendgroupMsg.Params.AutoEscape = false
		sendgroupMsg.Params.Message = sendMsg
		hub.Sendmsg <- sendgroupMsg
	} else {
		sendprivateMsg.Action = "send_private_msg"
		sendprivateMsg.Params.UserID = userId
		sendprivateMsg.Echo = ""
		sendprivateMsg.Params.AutoEscape = false
		sendprivateMsg.Params.Message = sendMsg
		hub.Sendmsg <- sendprivateMsg
	}

}
