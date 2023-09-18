package cover

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"os/exec"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 从视频链接中截取一帧并返回
func GetCoverFromUrl(url string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(url).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), nil
}

// 从视频字节中获取一帧返回
func GetCoverFromBytes(video []byte) ([]byte, error) {
	// 如果直接使用GetCoverFrom Url，需要等待storage刷新。
	// 而且这样还挺慢的。
	// 所以就有了这个函数。本来想使用pipe的，结果发现,
	// go的ffmpeg好像不支持使用管道作为输入。
	// 看到了stack overflow上建议使用
	// https://github.com/bakape/thumbnailer（需要pkg-config）
	// 不过要部署到1024code上，但是1024code无法安装pkg-config
	// 没办法，就使用exec.Command了。
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vframes", "1", "-f", "mjpeg", "-")
	cmd.Stdin = bytes.NewReader(video)

	var imageBuffer bytes.Buffer
	cmd.Stdout = &imageBuffer
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	return imageBuffer.Bytes(), nil
}

