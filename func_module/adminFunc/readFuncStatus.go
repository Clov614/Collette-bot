package adminFunc

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/setting"
	"regexp"
	"strconv"
)

func ReadfuncStatus(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData, check *FuncCheck) {
	receiveMsg := msgEvent.Message
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
		targets := []string{"功能状态", "功能列表"}
		for _, v := range targets {
			reRead, _ := regexp.Compile(v)
			if reRead.MatchString(receiveMsg) {
				dataCheck.Status = true
				dataCheck.SendMsg = printFunclist(check)
			}
		}

	}
}

// 将功能列表格式化输出
func printFunclist(check *FuncCheck) string {
	resultMsg := "\t\t\t\t功能列表  \n"
	for i, v := range check.FuncStatus {
		resultMsg += "功能:" + i + "\n"
		resultMsg += "状态:" + strconv.FormatBool(v) + "\n\n"
	}
	return resultMsg
}
