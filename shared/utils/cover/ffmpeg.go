package cover

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/bakape/thumbnailer"
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
	// go的ffmpeg-go好像不支持使用管道作为输入。
	// 我使用过管道、bytes.Reader，都不行，后面想到了他们不支持Seek。
	// 似乎ffmpeg需要对一些视频格式进行Seek，所以必须要一个支持它的东西。
	// 你可以试试用cat video.mp4 | ffmpeg -i ...
	// 这么做，一些视频文件会报错（应该是大多数文件，手头的mkv文件不会）
	// 但是ffmpeg -i ... < video.mp4不会，原因是这样实际上是打开了video.mp4，
	// 这样打开的，它支持Seek操作。
	// 但是实际上，ReadSeeker也不行，可能它只是对golang的内存读写做了封装了，
	// 而并没有对Linux进行封装。
	/*
	cmd := exec.Command("ffmpeg", "-i", "-", "-vframes", "1", "-f", "mjpeg", "-")
	var image bytes.Buffer
	cmd.Stdout = &image
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("Get cover failed: %w: %s", err, stderr.String())
		return nil, err
	}
	*/

	// 看到了stack overflow上建议使用
	// https://github.com/bakape/thumbnailer
	// 就决定用thumbnailer了。
	// 但是它有v1和v2两个版本
	// v1有一个好用的ProcessBuffer([]byte, options)
	// 但是v2没有，
	// 如果你要使用v1，你需要修改他们的源代码，修改这个文件
	// pkg/mod/github.com/bakape/thumbnailer@v1.0.0/ffmpeg.h
	// 加上这一行
	// #include <libavcodec/avcodec.h>
	//
	// 因为这个函数的定义是
	// int codec_context(AVCodecContext **avcc,
	//			  int *stream,
	//			  AVFormatContext *avfc,
	//			  const enum AVMediaType type);
	// 但是缺少该头文件，那么就缺少AVCodecContext这个定义

	thumbnailDimensions := thumbnailer.Dims{Width: 1080, Height: 2060}

	thumbnailOptions := thumbnailer.Options{JPEGQuality:100, MaxSourceDims:thumbnailer.Dims{}, ThumbDims:thumbnailDimensions, AcceptedMimeTypes: nil}

	_, thumbnail, err := thumbnailer.ProcessBuffer(video, thumbnailOptions)
	if err != nil {
		return nil, err
	}

	return thumbnail.Image.Data, nil
}

