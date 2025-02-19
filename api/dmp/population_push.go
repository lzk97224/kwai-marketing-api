package dmp

import (
	"github.com/lzk97224/kwai-marketing-api/core"
	"github.com/lzk97224/kwai-marketing-api/model/dmp"
)

// PopulationPush 人群包上线接口
func PopulationPush(clt *core.SDKClient, accessToken string, req *dmp.PopulationPushRequest) error {
	return clt.Post(accessToken, req, nil)
}
