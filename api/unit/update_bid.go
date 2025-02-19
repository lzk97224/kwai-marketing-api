package unit

import (
	"github.com/lzk97224/kwai-marketing-api/core"
	"github.com/lzk97224/kwai-marketing-api/model/unit"
)

// UpdateBid 修改广告组出价
func UpdateBid(clt *core.SDKClient, accessToken string, req *unit.UpdateBidRequest) error {
	return clt.Post(accessToken, req, nil)
}
