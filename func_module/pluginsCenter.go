package func_module

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/func_module/bilibiliAnalysis"
	"Collette_bot/func_module/ping"
)

func PluginsDetermine(msgEvent BaseEvent.GeneralMsg) (bool, string) {
	var checkData BaseEvent.PluginsData
	// 在此添加新的功能模块
	ping.Ping(msgEvent, &checkData)

	//log.Info("pluginsDetermine: ", receive)
	bilibiliAnalysis.BiliAnalysis(msgEvent, &checkData)

	bilibiliAnalysis.BilirawUrlanalysis(msgEvent, &checkData)
	// 向内层进行通信
	if checkData.Status {
		return true, checkData.SendMsg
	}
	return false, "PluginsDetermine_nil"
}
