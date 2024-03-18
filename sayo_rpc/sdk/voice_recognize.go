package sdk

import (
	"github.com/grteen/sayo_utils/constant"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
)

func PostVoiceRecognizeLocalFile(recognizeAddr, path string) (string, error) {
	result, err := sayorpc.DirectPost(recognizeAddr, constant.VoiceRecognizeURL, &VoiceRecognizeLocalFileReq{
		Path: path,
	})
	if err != nil {
		return "", err
	}

	return result.(string), nil
}
