package main

import (
	"fmt"
	"osu-ranked-downloder/osuapi"
)

func main() {
	fmt.Println("start")
	info := osuapi.GetMapsInfo("", "0")
	for _, beatmapSet := range info.BeatmapSets {
		fmt.Println(beatmapSet.Id, " ", beatmapSet.Title, " ranked: ", beatmapSet.Ranked == 1)
	}
}
