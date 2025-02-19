package target

import (
	"github.com/lzk97224/kwai-marketing-api/core"
	"github.com/lzk97224/kwai-marketing-api/model/target"
)

// TemplateDelete 删除定向模板
func TemplateDelete(clt *core.SDKClient, accessToken string, req *target.TemplateDeleteRequest) error {
	return clt.Post(accessToken, req, nil)
}
