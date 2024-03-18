package sayoinnerhttp

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
	"github.com/grteen/sayo_utils/utils"
)

func CallCoreToPullCenter(coreAddr string) error {
	url := utils.StringPlus("http://", coreAddr, constant.CorePullCenterURL)
	code, body, err := sayorpc.Post(url, struct{}{})
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrCallCoreToPullCenterFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrCallCoreToPullCenterFailed, int(resp.Code), resp.Msg)
	}

	return nil
}

func CoreVoiceCommand(coreAddr string, path string) error {
	url := utils.StringPlus("http://", coreAddr, constant.CoreVoiceCommand)
	code, body, err := sayorpc.Post(url, map[string]interface{}{constant.CoreVoiceCommandJSONPath: path})
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrCallCoreVoiceCommandFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrCallCoreVoiceCommandFailed, int(resp.Code), resp.Msg)
	}

	return nil
}
