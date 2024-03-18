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

type VoiceRecognizeLocalFileReq struct {
	Path string `json:"path"`
}

func PostVoiceRecognizeLocalFile(recognizeAddr, path string) (result string, err error) {
	code, body, err := sayorpc.Post(utils.StringPlus("http://", recognizeAddr, constant.VoiceRecognizeURL), &VoiceRecognizeLocalFileReq{path})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrVoiceRecognizeFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrVoiceRecognizeFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data.(string), nil
}
