package api

import (
	"context"
	"encoding/json"
	"reflect"
)

// SocialInfoList 是社交信息列表的兼容类型，可同时处理 JSON 数组和字符串。
// KOOK API 在社交信息为空时可能返回空字符串而非空数组。
type SocialInfoList []SocialInfo

// UnmarshalJSON 实现 json.Unmarshaler 接口，支持从数组或字符串反序列化。
func (s *SocialInfoList) UnmarshalJSON(data []byte) error {
	// 尝试作为数组解析
	var arr []SocialInfo
	if err := json.Unmarshal(data, &arr); err == nil {
		*s = arr
		return nil
	}
	// 尝试作为字符串解析（空字符串或无效格式时置为空数组）
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*s = nil
		return nil
	}
	return &json.UnmarshalTypeError{Value: string(data), Type: reflect.TypeOf(s)}
}

// IntimacyInfo 表示用户亲密度信息。
type IntimacyInfo struct {
	// ImgURL 是机器人给用户显示的形象图片地址。
	ImgURL string `json:"img_url"`
	// Social 是机器人显示给用户的社交信息。
	Social SocialInfoList `json:"social_info"`
	// LastRead 是用户上次查看的时间戳。
	LastRead int64 `json:"last_read"`
	// Score 是亲密度分数，范围 0-2200。
	Score int `json:"score"`
	// ImgList 是形象图片的总列表。
	ImgList []ImgInfo `json:"img_list"`
}

// ImgInfo 表示亲密度形象图片信息。
type ImgInfo struct {
	// ID 是形象图片的 ID。
	ID string `json:"id"`
	// URL 是形象图片的地址。
	URL string `json:"url"`
}

// SocialInfo 表示亲密度中的社交信息。
type SocialInfo struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// GetIntimacy 获取指定用户的亲密度信息。
// GET /intimacy/index?user_id={userID}
func GetIntimacy(ctx context.Context, client Doer, userID string) (*IntimacyInfo, error) {
	var result IntimacyInfo
	err := client.Do(ctx, "GET", "/intimacy/index", map[string]interface{}{
		"user_id": userID,
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateIntimacy 更新指定用户的亲密度信息。
// POST /intimacy/update
//
// 参数说明：
//   - userID: 目标用户 ID（必填）
//   - score: 亲密度分数，传 0 表示不修改
//   - socialInfo: 社交信息描述，为空表示不修改
//   - imgID: 亲密度图片 ID，传 0 表示不修改
func UpdateIntimacy(ctx context.Context, client Doer, userID string, score int, socialInfo string, imgID int) error {
	body := map[string]interface{}{
		"user_id": userID,
	}
	if score > 0 {
		body["score"] = score
	}
	if socialInfo != "" {
		body["social_info"] = socialInfo
	}
	if imgID > 0 {
		body["img_id"] = imgID
	}
	return client.Do(ctx, "POST", "/intimacy/update", body, nil)
}
