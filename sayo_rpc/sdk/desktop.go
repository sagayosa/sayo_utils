package sdk

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
	"github.com/grteen/sayo_utils/utils"
)

func GetWindow(desktopAddr string, way string, uuid string, argument map[string]interface{}) (result interface{}, err error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowExposeURL, "/", way, "/", uuid)
	code, body, err := sayorpc.Get(url, argument)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		err = sayoerror.ErrorInStatusCode(sayoerror.ErrGetWindowFailed, code)
		return
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		err = sayoerror.ErrorInMsgCode(sayoerror.ErrGetWindowFailed, int(resp.Code), resp.Msg)
		return
	}

	return resp.Data, nil
}

func PutWindow(desktopAddr string, way string, uuid string, argument interface{}) (result interface{}, err error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowExposeURL, "/", way, "/", uuid)
	code, body, err := sayorpc.Put(url, argument)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		err = sayoerror.ErrorInStatusCode(sayoerror.ErrPutWindowFailed, code)
		return
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		err = sayoerror.ErrorInMsgCode(sayoerror.ErrPutWindowFailed, int(resp.Code), resp.Msg)
		return
	}

	return resp.Data, nil
}
