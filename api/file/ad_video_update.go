package file

import (
	"github.com/lzk97224/kwai-marketing-api/core"
	"github.com/lzk97224/kwai-marketing-api/model/file"
)

// AdVideoUpdate 视频库-批量更新视频功能
func AdVideoUpdate(clt *core.SDKClient, accessToken string, req *file.AdVideoUpdateRequest) error {
	return clt.Post(accessToken, req, nil)
}
