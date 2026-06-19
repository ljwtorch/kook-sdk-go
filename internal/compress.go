package internal

import (
	"bytes"
	"compress/zlib"
	"io"
)

// Compress 使用 zlib 算法对数据进行压缩。
// 返回压缩后的字节切片和可能的错误。
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)

	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompress 使用 zlib 算法对数据进行解压。
// 返回解压后的字节切片和可能的错误。
func Decompress(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}
