package func_module

import (
	"Collette_bot/func_module/bilibiliAnalysis"
	"Collette_bot/func_module/ping"
)

func PluginsDetermine(receive string) (bool, string) {
	// 在此添加新的功能模块
	status, message := ping.Ping(receive)
	if status {
		return true, message
	}
	//log.Info("pluginsDetermine: ", receive)
	status, message = bilibiliAnalysis.BiliAnalysis(receive)
	if status {
		return true, message
	}

	status, message = bilibiliAnalysis.BilirawUrlanalysis(receive)
	// 向内层进行通信
	if status {
		return true, message
	}
	return false, "PluginsDetermine_nil"
}
