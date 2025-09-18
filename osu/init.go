package osu

import (
	"fmt"
	"net/http"
)

func getRequest(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("osuApiV1 failed, err:%v\n", err)
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", "https://osu.ppy.sh/beatmapsets")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36 Edg/140.0.0.0")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Cookie", fmt.Sprintf("osu_session=%s", osuSession))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("osuApiV1 failed, err:%v\n", err)
	}
	return resp
}
