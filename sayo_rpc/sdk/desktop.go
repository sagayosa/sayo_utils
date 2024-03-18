package sdk

import (
	"github.com/grteen/sayo_utils/constant"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
	"github.com/grteen/sayo_utils/utils"
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
