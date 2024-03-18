package proxy

import (
	"github.com/grteen/sayo_utils/constant"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
	"github.com/grteen/sayo_utils/sayo_rpc/sdk"
	"github.com/grteen/sayo_utils/utils"
)

func RegisterHotKey(frameworkAddr string, req *sdk.RegisterHotKeyReq) error {
	_, err := sayorpc.DirectPost(frameworkAddr, constant.ProxyDesktopRegisterHotKeyURL, req)
	if err != nil {
		return err
	}
	return nil
}

func OpenFileSelector(frameworkAddr string) (string, error) {
	result, err := sayorpc.DirectGet(frameworkAddr, constant.ProxyDesktopFileSelectorURL, map[string]interface{}{})
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func NewWindow(frameworkAddr string, req *sdk.NewWindowReq) (string, error) {
	result, err := sayorpc.DirectPost(frameworkAddr, constant.ProxyDesktopNewWindowURL, req)
	if err != nil {
		return "", err
	}
	return result.(string), err
}

func GetWindow(frameworkAddr string, way string, uuid string, argument map[string]interface{}) (interface{}, error) {
	return sayorpc.DirectGet(frameworkAddr, utils.StringPlus(constant.ProxyDesktopWindowExposeURL, "/", way, "/", uuid), argument)
}

func PutWindow(frameworkAddr string, way string, uuid string, argument interface{}) (result interface{}, err error) {
	return sayorpc.DirectPut(frameworkAddr, utils.StringPlus(constant.ProxyDesktopWindowExposeURL, "/", way, "/", uuid), argument)
}

func CursorPosition(frameworkAddr string) (x int, y int, err error) {
	result, err := sayorpc.DirectGet(frameworkAddr, constant.ProxyDesktopCursorPossition, map[string]interface{}{})
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

func WorkArea(frameworkAddr string) (width int, height int, err error) {
	result, err := sayorpc.DirectGet(frameworkAddr, constant.ProxyDesktopWorkArea, map[string]interface{}{})
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
