package Subs

import (
	"Collette_bot/network/ws"
	SendAPI "Collette_bot/sendMsgApi"
	"Collette_bot/setting"
	"fmt"
	"github.com/PassTheMayo/mcstatus/v3"
	log "github.com/sirupsen/logrus"
	"os"
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
type SubsMcSrvInfos struct {
	Each map[string]SubsMcSrvInfo `json:"each"`
}
type SubsMcSrvInfo struct {
	UserId   int      `json:"userId"`
	GroupID  int      `json:"groupID"`
	AddrSrv  string   `json:"addrSrv"`
	Interval int64    `json:"interval"`
	TempTime int64    `json:"tempTime"`
	TempInfo TempInfo `json:"temp_infos"`
}

type TempInfo struct {
	Players int `json:"players"`
}

var (
	sendgroupMsg   SendAPI.SENDGROUPMSG
	sendprivateMsg SendAPI.SENDPRIVATEMSG
	//pluginsMsg     BaseEvent.PluginsMsg
	//SubBaseinfo SubBaseInfo
	SubMC SubsMcSrvInfos
)

func init() {
	//SubBaseinfo = SubBaseInfo{
	//	StartTime: time.Now().Unix(),
	//	Interval:  60, // 默认订阅心跳为一分钟(节省性能开销)
	//}
	if !setting.PathExists("./Source/SubMC.yml") {
		if !setting.PathExists("./Source") {
			err := os.Mkdir("./Source", 0666)
			if err != nil {
				log.Error(err)
				os.Exit(0666)
			}
		}
		setting.WriteYaml(SubMC, "./Source/SubMC.yml")
	}
	setting.ReadYaml(&SubMC, "./Source/SubMC.yml")
}

// 处理订阅服务
func SubscribeHandle(hub *ws.Hub) {
	SubMC.SubsMcServerInfo(hub)
}

func (SubMc *SubsMcSrvInfos) SubsMcServerInfo(hub *ws.Hub) {
	// 读取
	setting.ReadYaml(&SubMC, "./Source/SubMC.yml")

	if len(SubMc.Each) == 0 {
		return
	}
	// 遍历每个id
	for i, v := range SubMc.Each {
		// 防止Players为空
		if strconv.Itoa(SubMc.Each[i].TempInfo.Players) == "" {
			v.TempInfo = TempInfo{
				Players: 0,
			}
		}
		// 验证是否满足间隔时间，执行订阅内容
		// 设置订阅更新间隔
		if (time.Now().Unix() - v.TempTime) >= v.Interval {
			v.TempTime = time.Now().Unix()
			// 获取到Mc服务器信息，以及当前玩家数量
			mcInfoMsg, nowPlayer, err := queryMcStatus(v.AddrSrv, v.TempInfo.Players)
			if err != nil {
				return
			}
			// 设置条件满足更新订阅
			if nowPlayer-v.TempInfo.Players != 0 {
				v.TempInfo = TempInfo{
					Players: nowPlayer,
				}
				SubMc.Each[i] = v
				setting.WriteYaml(&SubMc, "./Source/SubMC.yml")
				if i == (strconv.Itoa(v.UserId) + v.AddrSrv) {
					sendMsg("private", hub, mcInfoMsg, v.UserId, 0)
				} else {
					sendMsg("group", hub, mcInfoMsg, 0, v.GroupID)
				}
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
