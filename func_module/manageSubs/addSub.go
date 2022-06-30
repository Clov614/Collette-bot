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
					SubMC.Interval = int64(interval * 60)
				}
				SubMC.TempTime = time.Now().Unix()

			}
			if !existInSliceStr(SubMC.AddrSrv, addrSrv) {
				SubMC.AddrSrv = append(SubMC.AddrSrv, addrSrv)
			}
			if msgEvent.MessageType == "group" {
				if !existInSlice(SubMC.GroupIDs, msgEvent.GroupID) {
					SubMC.GroupIDs = append(SubMC.GroupIDs, msgEvent.GroupID)
				}
			} else {
				if !existInSlice(SubMC.UserIds, msgEvent.UserId) {
					SubMC.UserIds = append(SubMC.UserIds, msgEvent.UserId)
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
