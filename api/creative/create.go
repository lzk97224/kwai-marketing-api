package creative

import (
	"github.com/lzk97224/kwai-marketing-api/core"
	"github.com/lzk97224/kwai-marketing-api/model/creative"
)

// Create 创建创意
// 【注】联盟广告不支持便利贴图片素材，只有联盟广告支持横版竖版图片素材。
func Create(clt *core.SDKClient, accessToken string, req *creative.CreateRequest) (uint64, error) {
	var resp creative.CreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CreativeID, nil
}
