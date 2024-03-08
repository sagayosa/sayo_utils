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

func WindowHide(frameworkAddr string, uuid string) error {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopWindowHideURL)
	code, body, err := utils.Put(url, struct {
		UUID string `json:"uuid"`
	}{UUID: uuid})
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrWindowHideFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrWindowHideFailed, int(resp.Code), resp.Msg)
	}

	return nil
}

func WindowShow(frameworkAddr string, uuid string) error {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopWindowShowURL)
	code, body, err := utils.Put(url, struct {
		UUID string `json:"uuid"`
	}{UUID: uuid})
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrWindowShowFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrWindowShowFailed, int(resp.Code), resp.Msg)
	}

	return nil
}

func WindowSetPosition(frameworkAddr string, uuid string, x int, y int) error {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopWindowSetPosition)
	code, body, err := utils.Put(url, struct {
		UUID string `json:"uuid"`
		X    int    `json:"x"`
		Y    int    `json:"y"`
	}{UUID: uuid, X: x, Y: y})
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrWindowSetPositionFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrWindowSetPositionFailed, int(resp.Code), resp.Msg)
	}

	return nil
}

func CursorPosition(frameworkAddr string) (x int, y int, err error) {
	url := utils.StringPlus("http://", frameworkAddr, constant.ProxyDesktopCursorPossition)
	code, body, err := utils.Get(url, nil)
	if err != nil {
		return
	}

	if code != http.StatusOK {
		err = sayoerror.ErrorInStatusCode(sayoerror.ErrWindowSetPositionFailed, code)
		return
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		err = sayoerror.ErrorInMsgCode(sayoerror.ErrWindowSetPositionFailed, int(resp.Code), resp.Msg)
		return
	}

	result := &struct {
		X int `json:"x"`
		Y int `json:"y"`
	}{}

	if err = utils.UnMarshalUnknownAny(resp.Data, result); err != nil {
		return
	}

	return result.X, result.Y, nil
}
