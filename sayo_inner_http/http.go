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

func GetModuleByRole(frameworkAddr string, role string) (result interface{}, err error) {
	code, body, err := utils.Get(utils.StringPlus("http://", frameworkAddr, constant.GetModuleByRoleURL), map[string]interface{}{constant.GetModuleByRoleQueryRole: role})
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, sayoerror.Msg(sayoerror.ErrCoreGetRoleFailed, "StatusCode = %v", code)
	}

	resp := &baseresp.BaseResp{}
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	if resp.Code != sayoerror.SuccessCode {
		return nil, sayoerror.ErrorInMsgCode(sayoerror.ErrCoreGetRoleFailed, int(resp.Code), resp.Msg)
	}

	t := &struct {
		Modules []interface{} `json:"modules"`
	}{}
	if err := utils.UnMarshalUnknownAny(resp.Data, t); err != nil {
		return nil, err
	}

	return t.Modules, nil
}

func GetModuleVoiceRecognize(frameworkAddr string) (res []*module.Module, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RoleVoiceRecognize)
	if err != nil {
		return
	}

	res = []*module.Module{}
	if err = utils.UnMarshalUnknownAny(body, &res); err != nil {
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
	if err = utils.UnMarshalUnknownAny(body, &res); err != nil {
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
	if err = utils.UnMarshalUnknownAny(body, &res); err != nil {
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
	if err = utils.UnMarshalUnknownAny(body, &res); err != nil {
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
	if err = utils.UnMarshalUnknownAny(body, &res); err != nil {
		return
	}

	return res, nil
}
