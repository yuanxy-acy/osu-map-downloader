package osu

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"osu-map-downloader/config"
)

var osuSession = config.OsuSession

func DownloadMaps(sid string) {
	resp := getRequest(fmt.Sprintf("https://osu.ppy.sh/beatmapsets/%s/download", sid))
	var err error
	if resp.StatusCode == 302 {
		url := resp.Header.Get("Location")
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
		resp, err = http.Get(url)
		if err != nil {
			panic(err)
		}
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if err != nil {
		panic(err)
	}
	err = downloadProgress(config.SongsPath+params["filename"], resp)
	if err != nil {
		panic(err)
	}
}
