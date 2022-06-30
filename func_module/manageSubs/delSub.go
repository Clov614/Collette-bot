package manageSubs

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/Subs"
	"Collette_bot/setting"
	"regexp"
	"strings"
)

func DelSubMC(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData) {
	receiveMsg := msgEvent.Message
	if existInSlice(setting.Data.Admin, msgEvent.UserId) {
		reAdd, _ := regexp.Compile("删除订阅")
		SubMC := Subs.SubMC
		if msgEvent.MessageType == "group" {
			SubMC.UserIds = DeleteSlice(SubMC.UserIds, msgEvent.UserId)
		} else {
			SubMC.GroupIDs = DeleteSlice(SubMC.GroupIDs, msgEvent.GroupID)
		}

		if reAdd.MatchString(receiveMsg) {
			splitMsg := strings.Split(receiveMsg, " ")
			addrSrv := splitMsg[1]
			SubMC.AddrSrv = DeleteSliceStr(SubMC.AddrSrv, addrSrv)

			dataCheck.Status = true
			dataCheck.SendMsg = "删除订阅成功: " + addrSrv
			setting.WriteYaml(SubMC, "./Source/SubMC.yml")
		}
	}

}

// 删除切片中指定元素
func DeleteSlice(a []int, elem int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == elem {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

// 删除切片中指定元素 string
func DeleteSliceStr(a []string, elem string) []string {
	for i := 0; i < len(a); i++ {
		if a[i] == elem {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}
