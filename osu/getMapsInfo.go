package osu

import (
	"encoding/json"
	"fmt"
	"io"
)

type MapInfo struct {
	BeatmapSets []struct {
		Id         int    `json:"id"`
		Title      string `json:"title"`
		Ranked     int    `json:"ranked"`
		RankedDate string `json:"ranked_date"`
		covers     struct {
			List  string `json:"list"`
			Cover string `json:"cover"`
		}
	} `json:"beatmapsets"`
	CursorString string `json:"cursor_string"`
	Total        int    `json:"total"`
}

func GetMapsInfo(cursorString string, modeCode string) MapInfo {
	resp := getRequest(fmt.Sprintf("https://osu.ppy.sh/beatmapsets/search?m=%s&sort=ranked_asc&cursor_string=%s", modeCode, cursorString))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
	}
	info := MapInfo{}
	err = json.Unmarshal(b, &info)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
	}
	return info
}
