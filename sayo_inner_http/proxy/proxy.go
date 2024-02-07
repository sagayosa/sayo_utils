package proxy

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoinnerhttp "github.com/grteen/sayo_utils/sayo_inner_http"
	"github.com/grteen/sayo_utils/utils"
)

func PostAIDecisionRootCommand(frameworkAddr string, userCommand string) (result *sayoinnerhttp.AIDecisionResp, err error) {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyAICompletionsURL)
	code, body, err := utils.Post(url, map[string]interface{}{constant.ProxyAICompletionsJSONUserCommand: userCommand})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, sayoerror.ErrorInStatusCode(sayoerror.ErrCallCoreVoiceCommandFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return nil, sayoerror.ErrorInMsgCode(sayoerror.ErrCallCoreVoiceCommandFailed, int(resp.Code), resp.Msg)
	}

	result = &sayoinnerhttp.AIDecisionResp{}
	if err = utils.UnMarshalUnknownAny(resp.Data, result); err != nil {
		return
	}

	return result, nil
}
