package api

import (
	"context"
	"io"
)

// AssetResult 表示文件上传接口的响应结果。
type AssetResult struct {
	// URL 是上传成功后返回的资源访问地址。
	URL string `json:"url"`
}

// UploadAsset 上传文件/图片到 KOOK 平台。
// POST /asset/create (multipart/form-data)
//
// 参数说明：
//   - file: 文件内容的 io.Reader（必填）
//   - fileName: 文件名，用于设置 Content-Disposition（必填）
//
// 返回上传成功后的资源 URL，可在消息中引用。
func UploadAsset(ctx context.Context, client Doer, file io.Reader, fileName string) (*AssetResult, error) {
	var result AssetResult
	err := client.DoMultipart(ctx, "/asset/create", nil, "file", fileName, file, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
