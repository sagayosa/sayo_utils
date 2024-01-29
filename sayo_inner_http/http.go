package sayoinnerhttp

import (
	"encoding/json"
	"net/http"

	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	"github.com/grteen/sayo_utils/utils"
)

func GetModuleByRole(frameworkAddr string, role string) (body []byte, err error) {
	code, body, err := utils.Get(utils.StringPlus("http://", frameworkAddr, constant.GetModuleByRoleURL), map[string]interface{}{constant.GetModuleByRoleQueryRole: role})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, sayoerror.Msg(sayoerror.ErrCoreGetRoleFailed, "StatusCode = %v", code)
	}

	return body, nil
}

func GetModuleVoiceRecognize(frameworkAddr string) (res []*module.Module, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RoleVoiceRecognize)
	if err != nil {
		return
	}

	res = []*module.Module{}
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	return res, nil
}

func GetModuleVoiceGenerate(frameworkAddr string) (res []*module.Module, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RoleVoiceGenerate)
	if err != nil {
		return
	}

	res = []*module.Module{}
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	return res, nil
}

func GetModuleCore(frameworkAddr string) (res []*module.Module, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RoleCore)
	if err != nil {
		return
	}

	res = []*module.Module{}
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	return res, nil
}

func GetModuleAI(frameworkAddr string) (res []*module.Module, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RoleAI)
	if err != nil {
		return
	}

	res = []*module.Module{}
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	return res, nil
}

func GetModulePlugin(frameworkAddr string) (res []*module.Plugin, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RolePlugin)
	if err != nil {
		return
	}

	res = []*module.Plugin{}
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	return res, nil
}
