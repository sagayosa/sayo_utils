package sdk

import (
	"github.com/grteen/sayo_utils/constant"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
)

// func CallCoreToPullCenter(coreAddr string) error {
// 	url := utils.StringPlus("http://", coreAddr, constant.CorePullCenterURL)
// 	code, body, err := sayorpc.Post(url, struct{}{})
// 	if err != nil {
// 		return err
// 	}
// 	if code != http.StatusOK {
// 		return sayoerror.ErrorInStatusCode(sayoerror.ErrCallCoreToPullCenterFailed, code)
// 	}

// 	resp := &baseresp.BaseResp{}
// 	if err = json.Unmarshal(body, resp); err != nil {
// 		return err
// 	}
// 	if resp.Code != sayoerror.SuccessCode {
// 		return sayoerror.ErrorInMsgCode(sayoerror.ErrCallCoreToPullCenterFailed, int(resp.Code), resp.Msg)
// 	}

// 	return nil
// }

func CoreVoiceCommand(coreAddr string, path string) error {
	_, err := sayorpc.DirectPost(coreAddr, constant.CoreVoiceCommand, &CoreVoiceCommandReq{
		Path: path,
	})
	if err != nil {
		return err
	}

	return nil
}
