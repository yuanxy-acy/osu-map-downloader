package core

import (
	"encoding/json"
	"fmt"
	"osu-map-downloader/osu"
)

func FindRankedMaps() {
	info := osu.GetMapsInfo("", "0")
	for _, beatmapSet := range info.BeatmapSets {
		fmt.Println(fmt.Sprintf("%-14d", beatmapSet.Id), "ranked:", beatmapSet.Ranked == 1, "\t", beatmapSet.Title)
	}
	mapData := make(map[int]map[string]string)
	for len(info.BeatmapSets) > 0 && info.CursorString != "" {
		info = osu.GetMapsInfo(info.CursorString, "0")
		for _, beatmapSet := range info.BeatmapSets {
			fmt.Println(fmt.Sprintf("%-14d", beatmapSet.Id), "ranked:", beatmapSet.Ranked == 1, "\t", beatmapSet.Title)
			if beatmapSet.Ranked == 1 {
				mapData[beatmapSet.Id] = make(map[string]string)
				mapData[beatmapSet.Id]["title"] = beatmapSet.Title
				mapData[beatmapSet.Id]["ranked_date"] = beatmapSet.RankedDate
			}
		}
	}
	data, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err)
	}
	saveData("rankedMaps.json", &data)
}

func DownloadRankedMaps() {
	mapData := make(map[string]map[string]string)
	err := json.Unmarshal(readData("rankedMaps.json"), &mapData)
	if err != nil {
		fmt.Println(err)
		return
	}
	for sid, info := range mapData {
		fmt.Println("downloading  ", info["ranked_date"], "  ", fmt.Sprintf("%-15s", info["title"]))
		osu.DownloadMaps(sid)
	}
}
