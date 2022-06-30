package manageSubs

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/Subs"
	"Collette_bot/setting"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func AddSubMC(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData) {
	var Interval int64 = 60
	receiveMsg := msgEvent.Message
	if existInSlice(setting.Data.Admin, msgEvent.UserId) {
		reAdd, _ := regexp.Compile("添加订阅")
		SubMC := Subs.SubMC
		if reAdd.MatchString(receiveMsg) {
			splitMsg := strings.Split(receiveMsg, " ")
			addrSrv := splitMsg[1]
			if len(splitMsg) == 3 {
				interval, err := strconv.Atoi(splitMsg[2])
				if err != nil {
					log.Error(err)
					return
				}
				if interval >= 1 && interval <= 120 {
					Interval = int64(interval * 60)
				}
			}
			if msgEvent.MessageType == "group" {
				SubMC.Each[strconv.Itoa(msgEvent.GroupID)+addrSrv] = Subs.SubsMcSrvInfo{
					GroupID:  msgEvent.GroupID,
					UserId:   0,
					Interval: Interval,
					TempTime: time.Now().Unix(),
					AddrSrv:  addrSrv,
				}
			} else {
				SubMC.Each[strconv.Itoa(msgEvent.UserId)+addrSrv] = Subs.SubsMcSrvInfo{
					GroupID:  0,
					UserId:   msgEvent.UserId,
					Interval: Interval,
					TempTime: time.Now().Unix(),
					AddrSrv:  addrSrv,
				}
			}
			dataCheck.Status = true
			dataCheck.SendMsg = "添加订阅成功: " + addrSrv
			setting.WriteYaml(SubMC, "./Source/SubMC.yml")
		}

	}

}

// 判断元素是否存在于切片中
func existInSlice(a []int, elem int) (exist bool) {

	for _, v := range a {
		if v == elem {
			exist = true
			return exist
		}
	}
	exist = false
	return exist
}

// 判断元素是否存在于切片中
func existInSliceStr(a []string, elem string) (exist bool) {

	for _, v := range a {
		if v == elem {
			exist = true
			return exist
		}
	}
	exist = false
	return exist
}
