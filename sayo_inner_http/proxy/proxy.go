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

func PostAICompletion(frameworkAddr string, content string) (result interface{}, err error) {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyAICompletionsURL)
	code, body, err := utils.Post(url, map[string]interface{}{constant.ProxyAICompletionJSONContent: content})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, sayoerror.ErrorInStatusCode(sayoerror.ErrAIChatFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return nil, sayoerror.ErrorInMsgCode(sayoerror.ErrAIChatFailed, int(resp.Code), resp.Msg)
	}

	temp := &struct {
		Content string `json:"content"`
	}{}
	if err = utils.UnMarshalUnknownAny(resp.Data, temp); err != nil {
		return
	}

	return temp.Content, nil
}

func PostVoiceRecognizeLocalFile(frameworkAddr string, path string) (result string, err error) {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyVoiceRecognizeVoiceURL)
	code, body, err := utils.Post(url, map[string]interface{}{constant.ProxyVoiceRecognizeVoiceJSONPath: path})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrCallCoreVoiceCommandFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrCallCoreVoiceCommandFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data.(string), nil
}

func PostPlugin(frameworkAddr string, req *sayoinnerhttp.AIDecisionResp) error {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyPluginURL)
	code, body, err := utils.Post(url, req)
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
