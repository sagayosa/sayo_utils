package sdk

import (
	"github.com/sagayosa/sayo_utils/constant"
	"github.com/sagayosa/sayo_utils/module"
	sayorpc "github.com/sagayosa/sayo_utils/sayo_rpc"
	"github.com/sagayosa/sayo_utils/utils"
)

func GetWindow(desktopAddr string, way string, uuid string, argument map[string]interface{}) (result interface{}, err error) {
	result, err = sayorpc.DirectGet(desktopAddr, utils.StringPlus(constant.DesktopWindowExposeURL, "/", way, "/", uuid), argument)
	if err != nil {
		return
	}

	return result, nil
}

func PutWindow(desktopAddr string, way string, uuid string, argument interface{}) (result interface{}, err error) {
	result, err = sayorpc.DirectPut(desktopAddr, utils.StringPlus(constant.DesktopWindowExposeURL, "/", way, "/", uuid), argument)
	if err != nil {
		return
	}

	return result, nil
}

func NewWindow(desktopAddr string, req *NewWindowReq) (string, error) {
	result, err := sayorpc.DirectPost(desktopAddr, constant.DesktopNewWindowURL, req)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func OpenFileSelector(desktopAddr string) (string, error) {
	result, err := sayorpc.DirectGet(desktopAddr, constant.DesktopOpenFileSelectorURL, map[string]interface{}{})
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func RegisterHotKey(desktopAddr string, identifier string, hotkey *module.HotKey) error {
	_, err := sayorpc.DirectPost(desktopAddr, constant.DesktopRegisterHotKeyURL, &RegisterHotKeyReq{
		Identifier: identifier,
		Key:        hotkey.Key,
		Url:        hotkey.Url,
	})
	if err != nil {
		return err
	}
	return nil
}

func CursorPosition(desktopAddr string) (x int, y int, err error) {
	result, err := sayorpc.DirectGet(desktopAddr, constant.DesktopCursorPossition, map[string]interface{}{})
	if err != nil {
		return
	}

	resp := &struct {
		X int `json:"x"`
		Y int `json:"y"`
	}{}

	if err = utils.UnMarshalUnknownAny(result, resp); err != nil {
		return
	}

	return resp.X, resp.Y, nil
}

func WorkArea(desktopAddr string) (width int, height int, err error) {
	result, err := sayorpc.DirectGet(desktopAddr, constant.DesktopWorkArea, map[string]interface{}{})
	if err != nil {
		return
	}

	resp := &struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{}

	if err = utils.UnMarshalUnknownAny(result, resp); err != nil {
		return
	}

	return resp.Width, resp.Height, nil
}
