package main

import (
	"fmt"
	"osu-map-downloader/config"
	"osu-map-downloader/core"
)

func main() {
	fmt.Println("start")
	switch *config.Method {
	case "find":
		switch *config.MapType {
		case "ranked":
			core.FindRankedMaps()
		}
	case "download":
		switch *config.MapType {
		case "ranked":
			core.DownloadRankedMaps()
		}
	}
}
