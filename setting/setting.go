package setting

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

var (
	settingData SettingData
)

func init() {
	if !PathExists("./setting.yaml") {
		WriteYamlAppend(settingData, "./setting.yaml")
		log.Info("生成默认setting,yaml成功")
		log.Info("请配置setting.yaml后重启程序...")
		time.Sleep(time.Second * 5)
		log.Fatalln()
	}
	ReadYaml(&settingData, "./setting.yaml")
	log.Info(settingData.Nickname)
	if len(settingData.Nickname) == 0 {
		log.Info("未配置setting.yaml")
		log.Info("Nickname为空")
		time.Sleep(time.Second * 5)
		log.Fatalln()
	}

}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func WriteYaml(_type interface{}, path string) {
	//dataStr为[]byte,准备写入yaml
	dataStr, err := yaml.Marshal(_type)
	if err != nil {
		log.Fatalln("WriteYaml() Error: ", err)
	}

	err = ioutil.WriteFile(path, dataStr, 0644)
	if err != nil {
		log.Fatalln("WriteYaml() writeFile Error path: "+path, err)
	}
}

func ReadYaml(_type interface{}, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("读取Error path: "+path, err)
	}
	err = yaml.Unmarshal(file, _type)
	if err != nil {
		log.Fatalln("ERROR:"+path+" to data error: ", err)
	}
}

func WriteYamlAppend(_type interface{}, path string) {
	//dataStr为[]byte,准备写入yaml
	dataStr, err := yaml.Marshal(_type)
	if err != nil {
		log.Fatalln("WriteYaml() Error: ", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("WriteYaml() writeFile Error path: "+path, err)
	}
	helpStr := "# Nickname为机器人昵称\n# SelfQQ为机器人QQ号\n"
	file.Write([]byte(helpStr))
	file.Write(dataStr)
	file.Close()
}
