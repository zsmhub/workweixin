package apis

import (
	"bytes"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
)

const mediaFieldName = "media"

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeFile  = "file"
)

// 欲上传的素材
type Media struct {
	filename string
	filesize int64
	stream   io.Reader
}

func (m *Media) writeTo(w *multipart.Writer) error {
	wr, err := w.CreateFormFile(mediaFieldName, m.filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(wr, m.stream)
	if err != nil {
		return err
	}

	return nil
}

// NewMediaFromFile 从操作系统级文件创建一个欲上传的素材对象
func NewMediaFromFile(f *os.File) (*Media, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &Media{
		filename: stat.Name(),
		filesize: stat.Size(),
		stream:   f,
	}, nil
}

// NewMediaFromBuffer 从内存创建一个欲上传的素材对象
func NewMediaFromBuffer(filename string, buf []byte) (*Media, error) {
	stream := bytes.NewReader(buf)
	return &Media{
		filename: filename,
		filesize: int64(len(buf)),
		stream:   stream,
	}, nil
}

type UploadMediaReq struct {
	// 文件类型
	Type string `json:"type"`
	// 文件链接
	URL string `json:"url"`
}

type UploadMediaResult struct {
	// Type 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
	Type string `json:"type"`
	// MediaId 媒体文件上传后获取的唯一标识，3天内有效
	MediaId string `json:"media_id"`
	// CreatedAt 媒体文件上传时间戳
	CreatedAt string `json:"created_at"`
}

// UploadMedia 上传临时素材
func (c *ApiClient) UploadTempMedia(req UploadMediaReq) (UploadMediaResult, error) {
	var result UploadMediaResult

	// 下载文件
	_, body, err := FastClient.Get(nil, req.URL)
	if err != nil {
		return result, err
	}

	urlInfo, err := url.Parse(req.URL)
	if err != nil {
		return result, err
	}

	// 发现带上.gif后缀会导致前端无法引用到聊天窗口，故直接去掉后缀
	filename := uuid.New().String()
	if req.Type == "file" {
		filename = path.Base(urlInfo.Path)
	}

	media, err := NewMediaFromBuffer(filename, body)
	if err != nil {
		return result, err
	}

	uploadRes, err := c.ExecUploadMedia(ReqUploadMedia{
		Type:  req.Type,
		Media: media,
	})
	if err != nil {
		return result, err
	}

	result.MediaId = uploadRes.MediaId
	result.CreatedAt = uploadRes.CreatedAt
	result.Type = uploadRes.Type

	return result, nil
}
