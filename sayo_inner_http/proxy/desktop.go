package proxy

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	"github.com/grteen/sayo_utils/utils"
)

type RegisterHotKeyReq struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
	Key        string `json:"key"`
}

func RegisterHotKey(frameworkAddr string, req *RegisterHotKeyReq) error {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopRegisterHotKeyURL)
	code, body, err := utils.Post(url, req)
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrNewWindowFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrNewWindowFailed, int(resp.Code), resp.Msg)
	}

	return nil
}

func OpenFileSelector(frameworkAddr string) (result string, err error) {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopFileSelectorURL)
	code, body, err := utils.Get(url, nil)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrOpenFileSelectorFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrOpenFileSelectorFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data.(string), nil
}

type NewWindowReq struct {
	Theme  string      `json:"theme"`
	Url    string      `json:"url"`
	Frame  bool        `json:"frame"`
	Option interface{} `json:"option"`
}

func NewWindow(frameworkAddr string, req *NewWindowReq) (string, error) {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopNewWindowURL)
	code, body, err := utils.Post(url, req)
	if err != nil {
		return "", err
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrNewWindowFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return "", err
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrNewWindowFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data.(string), nil
}
