package osu

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Downloader struct {
	Resp       *http.Response
	Filename   string
	TotalSize  int64
	Downloaded int64
	StartTime  time.Time
	mu         sync.Mutex
}

func (d *Downloader) Download() error {
	d.StartTime = time.Now()

	file, err := os.OpenFile(d.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	resp := d.Resp

	// 获取文件大小
	if resp.Header.Get("Content-Length") != "" {
		d.TotalSize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			return err
		}
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("服务器错误: %s", resp.Status)
	}

	// 启动进度显示协程
	go d.showProgress()

	// 下载文件
	reader := &progressReader{
		reader: resp.Body,
		onRead: d.updateProgress,
	}

	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	fmt.Printf("\r")
	return nil
}

func (d *Downloader) updateProgress(n int64) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Downloaded += n
}

func (d *Downloader) showProgress() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		d.mu.Lock()
		downloaded := d.Downloaded
		total := d.TotalSize
		d.mu.Unlock()

		if downloaded >= total && total > 0 {
			return
		}

		if total <= 0 {
			fmt.Printf("\rDownloading... %s", formatBytes(downloaded))
			return
		}

		percent := float64(downloaded) / float64(total) * 100
		elapsed := time.Since(d.StartTime).Seconds()
		speed := float64(downloaded) / elapsed
		remaining := float64(total-downloaded) / speed

		fmt.Printf("\r%.1f%% | %s/%s | %.1f KB/s | ETA: %.1fs      ",
			percent,
			formatBytes(downloaded),
			formatBytes(total),
			speed/1024,
			remaining)
	}
}

type progressReader struct {
	reader io.Reader
	onRead func(int64)
}

func (r *progressReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if n > 0 {
		r.onRead(int64(n))
	}
	return n, err
}

func formatBytes(b int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	size := float64(b)
	i := 0
	for size >= 1024 && i < len(units)-1 {
		size /= 1024
		i++
	}
	return fmt.Sprintf("%.1f %s", size, units[i])
}

func downloadProgress(filepath string, resp *http.Response) error {

	downloader := &Downloader{
		Resp:     resp,
		Filename: filepath,
	}

	err := downloader.Download()
	return err
}
