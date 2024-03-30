package proxy

import (
	"github.com/sagayosa/sayo_utils/constant"
	sayorpc "github.com/sagayosa/sayo_utils/sayo_rpc"
	"github.com/sagayosa/sayo_utils/sayo_rpc/sdk"
	"github.com/sagayosa/sayo_utils/utils"
)

func PostAICompletion(frameworkAddr string, content string) (string, error) {
	result, err := sayorpc.DirectPost(frameworkAddr, constant.ProxyAICompletionsURL, &sdk.ProxyAICompletionReq{
		Content: content,
	})
	if err != nil {
		return "", err
	}

	resp := &struct {
		Content string `json:"content"`
	}{}
	if err = utils.UnMarshalUnknownAny(result, resp); err != nil {
		return "", err
	}

	return resp.Content, nil
}

func PostVoiceRecognizeLocalFile(frameworkAddr string, path string) (string, error) {
	result, err := sayorpc.DirectPost(frameworkAddr, constant.ProxyVoiceRecognizeVoiceURL, &sdk.ProxyVoiceRecognizeVoiceReq{
		Path: path,
	})
	if err != nil {
		return "", err
	}

	return result.(string), nil
}

func PostPlugin(frameworkAddr string, req *sdk.AIDecisionResp) error {
	_, err := sayorpc.DirectPost(frameworkAddr, constant.ProxyPluginURL, req)
	if err != nil {
		return err
	}
	return nil
}
