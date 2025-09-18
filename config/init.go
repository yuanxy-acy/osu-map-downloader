package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type conf struct {
	SongsPath  string `json:"songs_path"`
	OsuSession string `json:"osu_session"`
}

var (
	SongsPath  string
	Method     *string
	MapType    *string
	OsuSession string
)

func init() {
	configPath := flag.String("c", "config.json", "config file path")
	Method = flag.String("m", "none", "program running method")
	MapType = flag.String("t", "ranked", "program running type")
	flag.Parse()
	configJson, err := os.ReadFile(*configPath)
	if err != nil {
		if err.Error() == "open config.json: The system cannot find the file specified." {
			file, err := os.Create(*configPath)
			if err != nil {
				fmt.Println("create config.json file failed")
				fmt.Println(err)
				return
			}
			defer func(file *os.File) {
				err = file.Close()
				if err != nil {

				}
			}(file)
			_, err = file.WriteString("{\n    \"songs_path\": \"<your-path>\",\n    \"osu_session\": \"<your-osu-session-cookie-value>\"\n}")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("已创建默认配置文件")
			os.Exit(0)
		}
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
	dirInfo, err := os.ReadDir(config.SongsPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(dirInfo[0].Name())
	SongsPath = config.SongsPath
	OsuSession = config.OsuSession
}
