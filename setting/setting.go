package setting

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

var (
	Data SettingData
)

func init() {
	// 设置log格式
	log.SetFormatter(&nested.Formatter{
		NoColors:        true,
		ShowFullLevel:   true,
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
	})
	log.SetOutput(colorable.NewColorableStdout())

	if !PathExists("./setting.yml") {
		WriteYamlAppend(Data, "./setting.yml")
		log.Info("生成默认setting,yml成功")
		log.Info("请配置setting.yml后重启程序...")
		time.Sleep(time.Second * 5)
		os.Exit(0)
	}
	ReadYaml(&Data, "./setting.yml")
	if len(Data.Nickname) == 0 {
		log.Info("未配置setting.yml")
		log.Info("Nickname为空")
		time.Sleep(time.Second * 5)
		os.Exit(0)
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
		log.Error("WriteYaml() Error: ", err)
	}

	err = ioutil.WriteFile(path, dataStr, 0644)
	if err != nil {
		log.Error("WriteYaml() writeFile Error path: "+path, err)
	}
}

func ReadYaml(_type interface{}, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Error("读取Error path: "+path, err)
	}
	err = yaml.Unmarshal(file, _type)
	if err != nil {
		log.Error("ERROR:"+path+" to data error: ", err)
	}
}

func WriteYamlAppend(_type interface{}, path string) {
	//dataStr为[]byte,准备写入yaml
	dataStr, err := yaml.Marshal(_type)
	if err != nil {
		log.Error("WriteYaml() Error: ", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Error("WriteYaml() writeFile Error path: "+path, err)
	}
	helpStr := "# Nickname为机器人昵称\n# SelfQQ为机器人QQ号\n"
	file.Write([]byte(helpStr))
	file.Write(dataStr)
	file.Close()
}
