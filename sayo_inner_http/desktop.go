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

type NewWindowReq struct {
	Theme  string      `json:"theme"`
	Url    string      `json:"url"`
	Frame  bool        `json:"frame"`
	Option interface{} `json:"option"`
}

func NewWindow(desktopAddr string, req *NewWindowReq) (string, error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopNewWindowURL)
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

func GetWindow(desktopAddr string, way string, uuid string) (result interface{}, err error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowExposeURL, "/", way, "/", uuid)
	code, body, err := utils.Get(url, nil)
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
	code, body, err := utils.Put(url, argument)
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

func WindowHide(desktopAddr string, uuid string) error {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowHideURL)
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

func WindowShow(desktopAddr string, uuid string) error {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowShowURL)
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

func WindowSetPosition(desktopAddr string, uuid string, x int, y int) error {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowSetPosition)
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

func CursorPosition(desktopAddr string) (x int, y int, err error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopCursorPossition)
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

func WorkArea(desktopAddr string) (width int, height int, err error) {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWorkArea)
	code, body, err := utils.Get(url, nil)
	if err != nil {
		return
	}

	if code != http.StatusOK {
		err = sayoerror.ErrorInStatusCode(sayoerror.ErrWorkAreaFailed, code)
		return
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		err = sayoerror.ErrorInMsgCode(sayoerror.ErrWorkAreaFailed, int(resp.Code), resp.Msg)
		return
	}

	result := &struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{}

	if err = utils.UnMarshalUnknownAny(resp.Data, result); err != nil {
		return
	}

	return result.Width, result.Height, nil
}

func LoadURL(desktopAddr string, uuid string, targetUrl string) error {
	url := utils.StringPlus("http://", desktopAddr, constant.DesktopWindowURL)
	code, body, err := utils.Put(url, struct {
		UUID string `json:"uuid"`
		URL  string `json:"url"`
	}{UUID: uuid, URL: targetUrl})
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrLoadURLFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if resp.Code != sayoerror.SuccessCode {
		return sayoerror.ErrorInMsgCode(sayoerror.ErrLoadURLFailed, int(resp.Code), resp.Msg)
	}

	return nil
}
