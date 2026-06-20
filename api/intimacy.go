package api

import (
	"context"
)

// IntimacyInfo 表示用户亲密度信息。
type IntimacyInfo struct {
	// ImgURL 是机器人给用户显示的形象图片地址。
	ImgURL string `json:"img_url"`
	// SocialInfo 是机器人显示给用户的社交信息。
	SocialInfo string `json:"social_info"`
	// LastRead 是用户上次查看的时间戳。
	LastRead int64 `json:"last_read"`
	// LastModify 是最后修改时间戳。
	LastModify int64 `json:"last_modify"`
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

// GetIntimacy 获取指定用户的亲密度信息。
// GET /intimacy/index?user_id={userID}
//
// 参考文档：https://developer.kookapp.cn/doc/http/intimacy#获取用户亲密度
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
// 参考文档：https://developer.kookapp.cn/doc/http/intimacy#更新用户亲密度
//
// 参数说明：
//   - userID: 目标用户 ID（必填）
//   - score: 亲密度分数，传 0 表示不修改
//   - socialInfo: 社交信息描述，为空表示不修改
//   - imgID: 形象图片 ID，为空表示不修改
func UpdateIntimacy(ctx context.Context, client Doer, userID string, score int, socialInfo string, imgID string) error {
	body := map[string]interface{}{
		"user_id": userID,
	}
	if score > 0 {
		body["score"] = score
	}
	if socialInfo != "" {
		body["social_info"] = socialInfo
	}
	if imgID != "" {
		body["img_id"] = imgID
	}
	return client.Do(ctx, "POST", "/intimacy/update", body, nil)
}
