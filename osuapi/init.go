package osuapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"osu-ranked-downloder/config"
	"strings"
)

var accessToken string

var baseUrl = "https://osu.ppy.sh/api/v2/"
var contentType = "application/x-www-form-urlencoded"

func init() {
	data := "client_id=" + config.V2Api.ClientId + "&client_secret=" + config.V2Api.ClientSecret + "&grant_type=client_credentials&scope=public"
	resp, err := http.Post("https://osu.ppy.sh/oauth/token", contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("osuApiV1 failed, err:%v\n", err)
	}
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
	token := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}{}
	err = json.Unmarshal(b, &token)
	accessToken = token.AccessToken
}

type MapInfo struct {
	BeatmapSets []struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Ranked int    `json:"ranked"`
		covers struct {
			List  string `json:"list"`
			Cover string `json:"cover"`
		}
	} `json:"beatmapsets"`
	CursorString string `json:"cursor_string"`
	Total        int    `json:"total"`
}

func GetMapsInfo(cursorString string, modeCode string) MapInfo {
	req, err := http.NewRequest("GET", baseUrl+"beatmapsets/search?m="+modeCode+"&sort=ranked_asc&cursor_string="+cursorString, nil)
	if err != nil {
		fmt.Printf("osuApiV1 failed, err:%v\n", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("osuApiV1 failed, err:%v\n", err)
	}
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
