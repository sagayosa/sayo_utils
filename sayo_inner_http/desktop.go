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

func OpenFileSelector(desktopAddr string) (result string, err error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopOpenFileSelectorURL)
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

func RegisterHotKey(desktopAddr string, identifier string, hotkey *module.HotKey) error {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopRegisterHotKeyURL)
	code, body, err := utils.Post(url, struct {
		Identifier string `json:"identifier"`
		Key        string `json:"key"`
		Url        string `json:"url"`
	}{Identifier: identifier, Key: hotkey.Key, Url: hotkey.Url})
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrRegisterHotKeyFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrRegisterHotKeyFailed, int(resp.Code), resp.Msg)
	}

	return nil
}
