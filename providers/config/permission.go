package config

import "github.com/micro/go-config"

// Spec 验证
type Spec struct {
	Label       string `json:"label"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Config 获取配置
type Config struct {
	Validate map[string][]Spec `json:"Validate"`
}

// Conf 配置返回
var Conf Config

func init() {
	config.LoadFile("config/permission.json")
	config.Scan(&Conf)
}
