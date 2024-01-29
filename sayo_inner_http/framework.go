package sayoinnerhttp

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	"github.com/grteen/sayo_utils/utils"
)

func PullCenter(frameworkURL string) (result *module.Center, err error) {
	code, body, err := utils.Get(utils.StringPlus("http://", frameworkURL, constant.FrameWorkPullCenterURL), map[string]interface{}{})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, sayoerror.ErrorInStatusCode(sayoerror.ErrPullCenterFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return nil, sayoerror.ErrorInMsgCode(sayoerror.ErrPullCenterFailed, int(resp.Code), resp.Msg)
	}

	type temp struct {
		RoleMp map[string][]*module.Plugin `json:"role_map"`
	}
	t := &temp{RoleMp: make(map[string][]*module.Plugin)}
	if err := utils.UnMarshalUnknownAny(resp.Data, t); err != nil {
		return nil, err
	}

	result = module.NewCenter()
	for _, s := range t.RoleMp {
		for _, v := range s {
			if err := result.RegisterModule(v); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
