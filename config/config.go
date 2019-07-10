package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type HttpConfig struct {
	Port string `yaml:"port"`
}
type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
type ProjectConfig struct {
	HttpConfig  `yaml:"http"`
	MysqlConfig `yaml:"mysql"`
	Mode        string `yaml:"mode"`
	AppName     string `yaml:"app"`
}

var Config = &ProjectConfig{}

func init(){
	envFileName := "./config/config-local.yml"
	fileBytes, err := ioutil.ReadFile(envFileName)
	if err != nil {
		log.Fatalf("读取配置文件 %s 出错: %s", envFileName, err)
	}
	// 解析配置文件
	if err := yaml.Unmarshal(fileBytes, Config); err != nil {
		log.Fatalf("解析%s文件失败: %s", envFileName, err)
	}
}
