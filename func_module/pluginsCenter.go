package func_module

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/func_module/adminFunc"
	"Collette_bot/func_module/bilibiliAnalysis"
	"Collette_bot/func_module/manageSubs"
	"Collette_bot/func_module/ping"
	"Collette_bot/func_module/queryMCstatus"
	"Collette_bot/setting"
	log "github.com/sirupsen/logrus"
)

var (
	funcCheck adminFunc.FuncCheck = adminFunc.FuncCheck{map[string]bool{
		"ping":             true,
		"bilibiliAnalysis": true,
		"querymcStatus":    true,
	}}
)

func init() {
	if !setting.PathExists("./funcSetting.yml") {
		log.Info("正在初始化功能列表")
		setting.WriteYaml(funcCheck, "./funcSetting.yml")
		log.Info("初始化功能列表成功")
		log.Info("生成funcSetting.yml成功（功能开启关闭设置，每次重启生效）")
	}
	setting.ReadYaml(&funcCheck, "./funcSetting.yml")
	log.Info("读取功能列表成功")

}

func PluginsDetermine(msgEvent BaseEvent.PluginsMsg) (bool, string) {
	var checkData BaseEvent.PluginsData

	// 在此添加新的功能模块
	//ping.Ping(msgEvent, &checkData)
	checkFUNCstatus("ping", ping.Ping, msgEvent, &checkData)

	// 管理员管理功能开关功能
	adminFunc.SwitchFunc(msgEvent, &checkData, &funcCheck)
	adminFunc.ReadfuncStatus(msgEvent, &checkData, &funcCheck)
	//bilibiliAnalysis.BiliAnalysis(msgEvent, &checkData)
	//bilibiliAnalysis.BilirawUrlanalysis(msgEvent, &checkData)
	checkFUNCstatus("bilibiliAnalysis", bilibiliAnalysis.BiliAnalysis, msgEvent, &checkData)
	checkFUNCstatus("bilibiliAnalysis", bilibiliAnalysis.BilirawUrlanalysis, msgEvent, &checkData)

	//queryMCstatus.QuerymcStatus(msgEvent, &checkData)
	checkFUNCstatus("querymcStatus", queryMCstatus.QuerymcStatus, msgEvent, &checkData)

	// 订阅相关
	manageSubs.AddSubMC(msgEvent, &checkData)
	manageSubs.DelSubMC(msgEvent, &checkData)
	// 向内层进行通信
	if checkData.Status {
		return true, checkData.SendMsg
	}
	return false, "PluginsDetermine_nil"
}

// 判断功能开启或关闭
func checkFUNCstatus(funcName string, Func interface{}, msgEvent BaseEvent.PluginsMsg, checkData *BaseEvent.PluginsData) {
	if funcCheck.FuncStatus[funcName] {
		Func.(func(BaseEvent.PluginsMsg, *BaseEvent.PluginsData))(msgEvent, checkData)
	}
}
