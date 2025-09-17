package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type apiV2 struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type conf struct {
	V2Api     apiV2  `json:"v2_api"`
	SongsPath string `json:"songs_path"`
}

var (
	V2Api     apiV2
	SongsPath string
)

func init() {
	configPath := flag.String("c", "config.json", "config file path")
	flag.Parse()
	configJson, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Println("读取配置文件失败")
		fmt.Println(err)
		os.Exit(1)
	}
	var config = conf{}
	err = json.Unmarshal(configJson, &config)
	if err != nil {
		fmt.Println("解析配置文件失败")
		fmt.Println(err)
		os.Exit(1)
	}
	V2Api = config.V2Api
	SongsPath = config.SongsPath
}
