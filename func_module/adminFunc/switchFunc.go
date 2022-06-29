package adminFunc

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/setting"
	"regexp"
)

func SwitchFunc(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData, check *FuncCheck) {
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
		switchCheck(true, receiveMsg, dataCheck, check)
		switchCheck(false, receiveMsg, dataCheck, check)

	}
}

// 执行修改功能
func switchCheck(switchdirec bool, receiveMsg string, dataCheck *BaseEvent.PluginsData, check *FuncCheck) {
	var targets []string
	if switchdirec == true {
		targets = []string{"开启", "启用"}
	} else {
		targets = []string{"关闭", "禁用"}
	}
	for _, v := range targets {
		reOpen, _ := regexp.Compile(v)
		if reOpen.MatchString(receiveMsg) {
			// 开启某项功能
			//openMsg := "^"+ v + "(功能)$"
			reName, _ := regexp.Compile(v + "([\\s\\S]*?)" + "(功能)*$")
			if reName.MatchString(receiveMsg) {
				funcNames := reName.FindStringSubmatch(receiveMsg)
				funcName := funcNames[1]
				if isInCheck(funcName, check) {
					check.FuncStatus[funcNames[1]] = switchdirec
					dataCheck.Status = true
					dataCheck.SendMsg = "修改 " + funcNames[1] + " 成功"
					setting.WriteYaml(check, "./funcSetting.yml")
				} else {
					dataCheck.Status = true
					dataCheck.SendMsg = "未找到 " + funcNames[1] + " 该功能"
				}
			}
		}
	}

}

// 是否存在于FuncCheck(功能列表)中
func isInCheck(funcName string, check *FuncCheck) bool {
	_, ok := check.FuncStatus[funcName]
	if ok {
		return true
	}
	return false
}
