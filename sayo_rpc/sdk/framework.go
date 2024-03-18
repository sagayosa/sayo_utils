package sdk

import (
	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
	"github.com/grteen/sayo_utils/utils"
)

// func PullCenter(frameworkURL string) (result *module.Center, err error) {
// 	code, body, err := sayorpc.Get(utils.StringPlus("http://", frameworkURL, constant.FrameWorkPullCenterURL), map[string]interface{}{})
// 	if err != nil {
// 		return
// 	}
// 	if code != http.StatusOK {
// 		return nil, sayoerror.ErrorInStatusCode(sayoerror.ErrPullCenterFailed, code)
// 	}

// 	resp := &baseresp.BaseResp{}
// 	if err = json.Unmarshal(body, resp); err != nil {
// 		return
// 	}
// 	if resp.Code != sayoerror.SuccessCode {
// 		return nil, sayoerror.ErrorInMsgCode(sayoerror.ErrPullCenterFailed, int(resp.Code), resp.Msg)
// 	}

// 	type temp struct {
// 		RoleMp map[string][]*module.Module `json:"role_map"`
// 	}
// 	t := &temp{RoleMp: make(map[string][]*module.Module)}
// 	if err := utils.UnMarshalUnknownAny(resp.Data, t); err != nil {
// 		return nil, err
// 	}

// 	result = module.NewCenter()
// 	for _, s := range t.RoleMp {
// 		for _, v := range s {
// 			if err := result.RegisterModule(v); err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	return result, nil
// }

func GetModuleByRole(frameworkAddr string, role string) (result interface{}, err error) {
	result, err = sayorpc.DirectGet(frameworkAddr, constant.GetModuleByRoleURL, map[string]interface{}{
		"Role": role})
	if err != nil {
		return
	}
	t := &struct {
		Modules []interface{} `json:"modules"`
	}{}
	if err := utils.UnMarshalUnknownAny(result, t); err != nil {
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

func GetModulePlugin(frameworkAddr string) (res []*module.Module, err error) {
	body, err := GetModuleByRole(frameworkAddr, constant.RolePlugin)
	if err != nil {
		return
	}

	res = []*module.Module{}
	if err = utils.UnMarshalUnknownAny(body, &res); err != nil {
		return
	}

	return res, nil
}

// func Heart(addr string) (bool, error) {
// 	code, _, err := sayorpc.Get(utils.StringPlus("http://", addr, "/heart"), nil)
// 	if err != nil {
// 		return false, err
// 	}
// 	if code != http.StatusOK {
// 		return false, nil
// 	}

// 	return true, nil
// }
